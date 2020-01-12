package comb

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cszczepaniak/cribbage-scorer/cards"
	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		n   int
		exp int
	}{{
		n:   -1,
		exp: 1,
	}, {
		n:   0,
		exp: 1,
	}, {
		n:   1,
		exp: 1,
	}, {
		n:   2,
		exp: 2,
	}, {
		n:   3,
		exp: 6,
	}, {
		n:   4,
		exp: 24,
	}}
	for _, tc := range tests {
		res := Factorial(tc.n)
		assert.Equal(t, tc.exp, res)
	}
}
func TestNChooseK(t *testing.T) {
	tests := []struct {
		n   int
		k   int
		exp int
	}{{
		n:   3,
		k:   1,
		exp: 3,
	}, {
		n:   6,
		k:   6,
		exp: 1,
	}, {
		n:   6,
		k:   4,
		exp: 15,
	}, {
		n:   6,
		k:   3,
		exp: 20,
	}}
	for _, tc := range tests {
		res := Nchoosek(tc.n, tc.k)
		assert.Equal(t, tc.exp, res)
	}
}

func TestCombinations(t *testing.T) {
	tests := []struct {
		superset []string
		n        int
	}{{
		superset: []string{`ah`, `2h`, `3h`},
		n:        1,
	}, {
		superset: []string{`ah`, `2h`, `3h`, `4h`, `5s`},
		n:        4,
	}, {
		superset: []string{`ah`, `2h`, `3h`, `4h`, `5s`},
		n:        3,
	}, {
		superset: []string{`ah`, `2h`, `3h`, `4h`, `5s`, `10d`},
		n:        3,
	}}
	for _, tc := range tests {
		cds := make([]cards.Card, len(tc.superset))
		for i, s := range tc.superset {
			c, err := cards.NewCardFromString(s)
			require.NoError(t, err)
			cds[i] = c
		}
		combs := Combinations(cds, tc.n)
		assert.Len(t, combs, Nchoosek(len(tc.superset), tc.n))
	}
}
