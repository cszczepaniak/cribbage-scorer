package comb

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cszczepaniak/cribbage-scorer/cards"
	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	tests := []struct {
		superset []string
		n        int
		expLen   int
	}{{
		superset: []string{`ah`, `2h`, `3h`},
		n:        1,
		expLen:   3,
	}, {
		superset: []string{`ah`, `2h`, `3h`, `4h`, `5s`},
		n:        4,
		expLen:   5,
	}, {
		superset: []string{`ah`, `2h`, `3h`, `4h`, `5s`},
		n:        3,
		expLen:   10,
	}, {
		superset: []string{`ah`, `2h`, `3h`, `4h`, `5s`, `10d`},
		n:        3,
		expLen:   20,
	}}
	for _, tc := range tests {
		cds := make([]cards.Card, len(tc.superset))
		for i, s := range tc.superset {
			c, err := cards.FromString(s)
			require.NoError(t, err)
			cds[i] = c
		}
		combs := Combinations(cds, tc.n)
		assert.Len(t, combs, tc.expLen)
	}
}
