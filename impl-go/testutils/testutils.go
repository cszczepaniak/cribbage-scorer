// +build !prod

package testutils

import (
	"testing"

	"github.com/cszczepaniak/cribbage-scorer/cards"
	"github.com/stretchr/testify/require"
)

func MakeHandAndCut(tb testing.TB, handStrs []string, cutStr string) ([]cards.Card, cards.Card) {
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

func ParseCards(tb testing.TB, strs []string) []cards.Card {
	cs := make([]cards.Card, len(strs))
	for i, s := range strs {
		c, err := cards.FromString(s)
		require.NoError(tb, err)
		cs[i] = c
	}
	return cs
}
