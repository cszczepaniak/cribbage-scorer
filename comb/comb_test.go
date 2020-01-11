package comb

import (
	"testing"

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
	type dummyType struct {
		a int
	}
	tests := []struct {
		superset []interface{}
		n        int
	}{{
		superset: []interface{}{1, 2, 3},
		n:        1,
	}, {
		superset: []interface{}{1, 2, 3, 4, 5},
		n:        3,
	}, {
		superset: []interface{}{1, 2, 3, 4, 5},
		n:        3,
	}, {
		superset: []interface{}{
			dummyType{a: 1},
			dummyType{a: 2},
			dummyType{a: 3},
			dummyType{a: 4},
			dummyType{a: 5},
			dummyType{a: 6},
		},
		n: 3,
	}}
	for _, tc := range tests {
		combs := Combinations(tc.superset, tc.n)
		assert.Len(t, combs, Nchoosek(len(tc.superset), tc.n))
	}
}
