package score

import (
	"errors"
	"time"

	"github.com/cszczepaniak/cribbage-scorer/cards"
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
		return scoreHandSerial(hand, cut, isCrib), nil
	}
	return scoreHandParallel(hand, cut, isCrib)
}

func scoreHandSerial(hand []cards.Card, cut cards.Card, isCrib bool) int {
	allCards := append(hand, cut)
	rankCounts := make(map[int]int, len(allCards))
	suits := make(map[cards.Suit]int, 4)
	values := make([]int, len(allCards))
	score := 0
	for i, c := range allCards {
		rankCounts[c.Rank]++
		suits[c.Suit]++
		values[i] = c.Value()
	}
	rs := &runScorer{}
	for _, c := range allCards {
		rs.scoreRunsIter(c, rankCounts)
		// score nobs while we're looping
		if c != cut && c.Rank == 11 && c.Suit == cut.Suit {
			score += 1
		}
	}
	score += rs.maxRunScore
	// flush
	if len(suits) == 1 {
		score += 5
	}
	if len(suits) == 2 && suits[cut.Suit] == 1 && !isCrib {
		score += 4
	}
	// fifteens
	vs := [5]int{
		hand[0].Value(),
		hand[1].Value(),
		hand[2].Value(),
		hand[3].Value(),
		cut.Value(),
	}
	score += scoreFifteensFromValueList(vs)
	// pairs
	score += scorePairsFromMap(rankCounts)

	return score
}

func scoreHandParallel(hand []cards.Card, cut cards.Card, isCrib bool) (int, error) {
	scores := make(chan int)
	go func() {
		scores <- scoreFifteens(hand, cut)
	}()
	go func() {
		scores <- scorePairs(hand, cut)
	}()
	go func() {
		scores <- scoreFlush(hand, cut, isCrib)
	}()
	go func() {
		scores <- scoreRuns(hand, cut)
	}()
	go func() {
		scores <- scoreNobs(hand, cut)
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

func scoreRuns(hand []cards.Card, cut cards.Card) int {
	all := append(hand, cut)
	rankCounts := make(map[int]int, len(all))
	for _, c := range all {
		rankCounts[c.Rank]++
	}
	mostPoints := 0
	var mult int
	for _, c := range all {
		if _, ok := rankCounts[c.Rank-1]; ok {
			// this is already part of a run; skip calculation
			continue
		}
		runLen := 1
		// we're at the potential beginning of a run
		nextUp := c.Rank + 1
		mult = rankCounts[c.Rank]
		for ct, ok := rankCounts[nextUp]; ok; ct, ok = rankCounts[nextUp] {
			mult *= ct
			runLen++
			nextUp++
		}
		if runLen >= 3 && runLen*mult > mostPoints {
			mostPoints = runLen * mult
		}
		mult = 1
		runLen = 1
	}
	return mostPoints
}

type runScorer struct {
	currRunLen  int
	currRunMult int
	maxRunScore int
}

func (rs *runScorer) scoreRunsIter(curr cards.Card, rankCounts map[int]int) {
	if _, ok := rankCounts[curr.Rank-1]; ok {
		// this is already part of a previously calculated run; skip calculation
		return
	}
	rs.currRunLen = 1
	rs.currRunMult = 1
	// we're at the potential beginning of a run
	nextUp := curr.Rank + 1
	rs.currRunMult = rankCounts[curr.Rank]
	for ct, ok := rankCounts[nextUp]; ok; ct, ok = rankCounts[nextUp] {
		rs.currRunMult *= ct
		rs.currRunLen++
		nextUp++
	}
	if rs.currRunLen >= 3 && rs.currRunLen*rs.currRunMult > rs.maxRunScore {
		rs.maxRunScore = rs.currRunLen * rs.currRunMult
	}
}

func scoreNobs(hand []cards.Card, cut cards.Card) int {
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

func scoreFlush(hand []cards.Card, cut cards.Card, isCrib bool) int {
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

func scorePairs(hand []cards.Card, cut cards.Card) int {
	allCards := append(hand, cut)
	rankCounts := make(map[int]int, len(allCards))
	for _, c := range allCards {
		rankCounts[c.Rank]++
	}
	return scorePairsFromMap(rankCounts)
}

func scorePairsFromMap(rankCounts map[int]int) int {
	score := 0
	for _, val := range rankCounts {
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

func scoreFifteens(hand []cards.Card, cut cards.Card) int {
	allCards := append(hand, cut)
	values := [5]int{
		hand[0].Value(),
		hand[1].Value(),
		hand[2].Value(),
		hand[3].Value(),
		cut.Value(),
	}
	for i, c := range allCards {
		values[i] = c.Value()
	}
	return scoreFifteensFromValueList(values)
}

func scoreFifteensFromValueList(values [5]int) int {
	if (values[0]|values[1]|values[2]|values[3]|values[4])&1 == 0 {
		return 0
	}
	sum := values[0] + values[1] + values[2] + values[3] + values[4]
	if sum < 15 || sum > 46 {
		return 0
	}
	return fifteens(0, values[:]...) * 2
}

func fifteens(sum int, hand ...int) int {
	if sum == 15 {
		return 1
	}
	if sum > 15 {
		return 0
	}

	var res int
	for i, c := range hand {
		res += fifteens(sum+c, hand[i+1:]...)
	}
	return res
}

func (s *Scorer) validateHand(hand []cards.Card) error {
	if len(hand) != 4 {
		return ErrInvalidHandSize
	}
	return nil
}
