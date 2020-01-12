package testutils

import "github.com/cszczepaniak/cribbage-scorer/cards"

func MakeHandAndCut(handStrs []string, cutStr string) ([]cards.Card, cards.Card, error) {
	hand := make([]cards.Card, len(handStrs))
	for i, s := range handStrs {
		c, err := cards.NewCardFromString(s)
		if err != nil {
			return nil, cards.Card{}, err
		}
		hand[i] = c
	}
	cut, err := cards.NewCardFromString(cutStr)
	if err != nil {
		return nil, cards.Card{}, err
	}
	return hand, cut, nil
}
