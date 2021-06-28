package all

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmitCombinations(t *testing.T) {
	// tests := []struct {
	// 	n   int
	// 	max int
	// 	exp [][]cards.Card
	// }{{
	// 	n:   2,
	// 	max: 1,
	// 	exp: [][]cards.Card{
	// 		[]cards.Card{{Suit: cards.Suit(0), Rank: 0}, {Suit: cards.Suit(0), Rank: 1}},
	// 		[]cards.Card{{Suit: cards.Suit(0), Rank: 1}, {Suit: cards.Suit(0), Rank: 1}},
	// 		[]cards.Card{{Suit: cards.Suit(1), Rank: 0}, {Suit: cards.Suit(0), Rank: 1}},
	// 		[]cards.Card{{Suit: cards.Suit(0), Rank: 0}, {Suit: cards.Suit(0), Rank: 1}},
	// 		[]cards.Card{{Suit: cards.Suit(0), Rank: 0}, {Suit: cards.Suit(0), Rank: 1}},
	// 		[]cards.Card{{Suit: cards.Suit(0), Rank: 0}, {Suit: cards.Suit(0), Rank: 1}},
	// 	},
	// }}
	// for _, tc := range tests {
	// 	res := make([]cards.Card, 0, len(tc.exp))
	// 	hChan := make(chan []cards.Card)
	// 	errChan := make(chan error)
	// 	go EmitCombinations(tc.n, tc.max, hChan, errChan)
	// 	for i := 0; i < len(tc.exp); i++ {
	// 		select {
	// 		case h := <-hChan:
	// 			res = append(res, h)
	// 		case err := <-errChan:
	// 		}
	// 	}
	// }
}

func TestIncState(t *testing.T) {
	tests := []struct {
		state    []int
		maxState int
		expState []int
	}{{
		state:    []int{0, 0, 0, 0, 0},
		maxState: 12,
		expState: []int{1, 0, 0, 0, 0},
	}, {
		state:    []int{12, 0, 0, 0, 0},
		maxState: 12,
		expState: []int{0, 1, 0, 0, 0},
	}, {
		state:    []int{12, 12, 0, 0, 0},
		maxState: 12,
		expState: []int{0, 0, 1, 0, 0},
	}, {
		state:    []int{12, 12, 12, 12, 12},
		maxState: 12,
		expState: []int{0, 0, 0, 0, 13},
	}, {
		state:    []int{51, 0, 0, 0, 0},
		maxState: 51,
		expState: []int{0, 1, 0, 0, 0},
	}}
	for _, tc := range tests {
		res := incState(tc.state, tc.maxState, 0)
		assert.Equal(t, tc.expState, res)
	}
}
