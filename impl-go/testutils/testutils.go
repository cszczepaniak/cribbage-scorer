// +build !prod

package testutils

import (
	"github.com/cszczepaniak/cribbage-scorer/cards"
	"github.com/stretchr/testify/require"
)

func MakeHandAndCut(tb require.TestingT, handStrs []string, cutStr string) ([]cards.Card, cards.Card) {
	hand := make([]cards.Card, len(handStrs))
	for i, s := range handStrs {
		c, err := cards.FromString(s)
		require.NoError(tb, err)
		hand[i] = c
	}
	cut, err := cards.FromString(cutStr)
	require.NoError(tb, err)
	return hand, cut
}
