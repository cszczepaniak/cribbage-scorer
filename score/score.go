package score

import (
	"errors"

	"github.com/cszczepaniak/cribbage-scorer/cards"
	"github.com/cszczepaniak/cribbage-scorer/comb"
)

var (
	ErrInvalidHandSize = errors.New(`a hand must have exactly four cards in it`)
)

func scorePairs(hand []cards.Card, cut cards.Card) (int, error) {
	err := validateHand(hand)
	if err != nil {
		return 0, err
	}
	all := append(hand, cut)
	combs := comb.Combinations(all, 2)
	score := 0
	for _, comb := range combs {
		if comb[0].Rank == comb[1].Rank {
			score += 2
		}
	}
	return score, nil
}

func scoreFifteens(hand []cards.Card, cut cards.Card) (int, error) {
	err := validateHand(hand)
	if err != nil {
		return 0, err
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
	return score, nil
}

func validateHand(hand []cards.Card) error {
	if len(hand) != 4 {
		return ErrInvalidHandSize
	}
	return nil
}
