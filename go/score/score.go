package score

import (
	"errors"
	"time"

	"github.com/cszczepaniak/cribbage-scorer/cards"
	"github.com/cszczepaniak/cribbage-scorer/comb"
)

var (
	ErrInvalidHandSize = errors.New(`a hand must have exactly four cards in it`)
)

type Scorer struct {
	isSerial bool
}

func NewSerialScorer() *Scorer {
	return &Scorer{
		isSerial: true,
	}
}

func NewParallelScorer() *Scorer {
	return &Scorer{
		isSerial: false,
	}
}

func (s *Scorer) ScoreHand(hand []cards.Card, cut cards.Card, isCrib bool) (int, error) {
	if err := s.validateHand(hand); err != nil {
		return 0, err
	}
	if s.isSerial {
		return s.scoreHandSerial(hand, cut, isCrib), nil
	}
	return s.scoreHandParallel(hand, cut, isCrib)
}

func (s *Scorer) scoreHandSerial(hand []cards.Card, cut cards.Card, isCrib bool) int {
	return int(s.scoreFifteens(hand, cut) + s.scorePairs(hand, cut) + s.scoreFlush(hand, cut, isCrib) +
		s.scoreRuns(hand, cut) + s.scoreNobs(hand, cut))
}

func (s *Scorer) scoreHandParallel(hand []cards.Card, cut cards.Card, isCrib bool) (int, error) {
	scores := make(chan int)
	go func() {
		scores <- s.scoreFifteens(hand, cut)
	}()
	go func() {
		scores <- s.scorePairs(hand, cut)
	}()
	go func() {
		scores <- s.scoreFlush(hand, cut, isCrib)
	}()
	go func() {
		scores <- s.scoreRuns(hand, cut)
	}()
	go func() {
		scores <- s.scoreNobs(hand, cut)
	}()
	total := 0
	for i := 0; i < 5; i++ {
		select {
		case s := <-scores:
			total += s
		case <-time.After(time.Second):
			return 0, errors.New(`parallel scoring timed out`)
		}
	}
	return total, nil
}

func (s *Scorer) scoreRuns(hand []cards.Card, cut cards.Card) int {
	all := append(hand, cut)
	for i := 5; i > 2; i-- {
		score := 0
		combs := comb.Combinations(all, i)
		for _, comb := range combs {
			score += s.scoreRun(comb)
		}
		if score > 0 {
			return score
		}
	}
	return 0
}

func (s *Scorer) scoreRun(set []cards.Card) int {
	sorted := cards.SortByRankAscending(set)
	for i := 0; i < len(sorted)-1; i++ {
		thisCard := sorted[i]
		nextCard := sorted[i+1]
		if thisCard.Rank+1 != nextCard.Rank {
			return 0
		}
	}
	return len(set)
}

func (s *Scorer) scoreNobs(hand []cards.Card, cut cards.Card) int {
	if cut.Rank == 11 {
		return 0
	}
	for _, c := range hand {
		if c.Rank == 11 && c.Suit == cut.Suit {
			return 1
		}
	}
	return 0
}

func (s *Scorer) scoreFlush(hand []cards.Card, cut cards.Card, isCrib bool) int {
	suitMap := make(map[cards.Suit]struct{}, len(hand)+1)
	suitMap[hand[0].Suit] = struct{}{}
	for _, c := range hand[1:] {
		if _, ok := suitMap[c.Suit]; !ok {
			return 0
		}
	}
	if _, ok := suitMap[cut.Suit]; !ok {
		if isCrib {
			return 0
		}
		return 4
	}
	return 5
}

func (s *Scorer) scorePairs(hand []cards.Card, cut cards.Card) int {
	all := append(hand, cut)
	m := make(map[int]int, len(all))
	for _, c := range all {
		m[c.Rank]++
	}
	score := 0
	for _, val := range m {
		switch val {
		case 2:
			score += 2
		case 3:
			score += 6
		case 4:
			score += 12
		}
	}
	return score
}

func (s *Scorer) scoreFifteens(hand []cards.Card, cut cards.Card) int {
	all := append(hand, cut)
	vals := make([]int, len(all))
	for i, c := range all {
		vals[i] = c.Value()
	}
	nFifteens := fifteens(0, vals...)
	return nFifteens * 2
}

func fifteens(sum int, hand ...int) int {
	if sum == 15 {
		return 1
	}
	if sum > 15 {
		return 0
	}

	switch len(hand) {
	case 1:
		return fifteens(sum + hand[0])
	case 2:
		return fifteens(sum+hand[0]) + fifteens(sum+hand[1]) + fifteens(sum+hand[0]+hand[1])
	case 3:
		return fifteens(sum+hand[0], hand[1], hand[2]) + fifteens(sum+hand[1], hand[2]) + fifteens(sum+hand[2])
	case 4:
		return fifteens(sum+hand[0], hand[1], hand[2], hand[3]) +
			fifteens(sum+hand[1], hand[2], hand[3]) +
			fifteens(sum+hand[2], hand[3]) +
			fifteens(sum+hand[3])
	case 5:
		return fifteens(sum+hand[0], hand[1], hand[2], hand[3], hand[4]) +
			fifteens(sum+hand[1], hand[2], hand[3], hand[4]) +
			fifteens(sum+hand[2], hand[3], hand[4]) +
			fifteens(sum+hand[3], hand[4]) +
			fifteens(sum+hand[4])
	default:
		return 0
	}
}

func (s *Scorer) validateHand(hand []cards.Card) error {
	if len(hand) != 4 {
		return ErrInvalidHandSize
	}
	return nil
}
