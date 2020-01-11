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
