package score

import (
	"errors"

	"github.com/cszczepaniak/cribbage-scorer/cards"
	"github.com/cszczepaniak/cribbage-scorer/comb"
)

var (
	ErrInvalidHandSize = errors.New(`a hand must have exactly four cards in it`)
)

func ScoreHand(hand []cards.Card, cut cards.Card, isCrib bool) (int, error) {
	err := validateHand(hand)
	if err != nil {
		return 0, err
	}
	return (scoreFifteens(hand, cut) + scorePairs(hand, cut) + scoreFlush(hand, cut, isCrib) +
		scoreRuns(hand, cut) + scoreNobs(hand, cut)), nil
}

func scoreRuns(hand []cards.Card, cut cards.Card) int {
	all := append(hand, cut)
	for i := 5; i > 2; i-- {
		score := 0
		combs := comb.Combinations(all, i)
		for _, comb := range combs {
			score += scoreRun(comb)
		}
		if score > 0 {
			return score
		}
	}
	return 0
}

func scoreRun(set []cards.Card) int {
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
	if isCrib {
		return scoreCribFlush(hand, cut)
	}
	return scoreHandFlush(hand, cut)
}

func scoreCribFlush(hand []cards.Card, cut cards.Card) int {
	suitMap := make(map[cards.Suit]struct{}, len(hand)+1)
	suitMap[hand[0].Suit] = struct{}{}
	for _, c := range hand[1:] {
		if _, ok := suitMap[c.Suit]; !ok {
			return 0
		}
	}
	if _, ok := suitMap[cut.Suit]; !ok {
		return 0
	}
	return 5
}

func scoreHandFlush(hand []cards.Card, cut cards.Card) int {
	suitMap := make(map[cards.Suit]struct{}, len(hand)+1)
	suitMap[hand[0].Suit] = struct{}{}
	for _, c := range hand[1:] {
		if _, ok := suitMap[c.Suit]; !ok {
			return 0
		}
	}
	score := 4
	if _, ok := suitMap[cut.Suit]; ok {
		score++
	}
	return score
}

func scorePairs(hand []cards.Card, cut cards.Card) int {
	err := validateHand(hand)
	if err != nil {
		return 0
	}
	all := append(hand, cut)
	combs := comb.Combinations(all, 2)
	score := 0
	for _, comb := range combs {
		if comb[0].Rank == comb[1].Rank {
			score += 2
		}
	}
	return score
}

func scoreFifteens(hand []cards.Card, cut cards.Card) int {
	err := validateHand(hand)
	if err != nil {
		return 0
	}
	all := append(hand, cut)
	score := 0
	for i := 2; i < 6; i++ {
		combs := comb.Combinations(all, i)
		for _, comb := range combs {
			val := 0
			for _, c := range comb {
				val += c.Value()
			}
			if val == 15 {
				score += 2
			}
		}
	}
	return score
}

func validateHand(hand []cards.Card) error {
	if len(hand) != 4 {
		return ErrInvalidHandSize
	}
	return nil
}
