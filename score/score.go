package score

import (
	"errors"

	"github.com/cszczepaniak/cribbage-scorer/comb"

	"github.com/cszczepaniak/cribbage-scorer/cards"
)

var (
	ErrInvalidHandSize = errors.New(`a hand must have exactly four cards in it`)
)

func scoreFifteens(hand []cards.Card, cut cards.Card, isCrib bool) (int, error) {
	err := validateHand(hand)
	if err != nil {
		return 0, err
	}
	all := append(hand, cut)
	score := 0
	for i := 2; i < 6; i++ {
		combs := comb.Combinations(hand, i)
		for _, comb := range combs {
			val := 0
			for _, c := range comb {
				card, ok := c.(cards.Card)
				if !ok {
					return 0, errors.New(`type assertion failed`)
				}
				val += card.Value()
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
