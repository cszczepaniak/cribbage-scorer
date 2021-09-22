package comb

import (
	"github.com/cszczepaniak/cribbage-scorer/cards"
)

func Combinations(superset []cards.Card, n int) [][]cards.Card {
	if len(superset) == n {
		return [][]cards.Card{superset}
	}
	if n == 1 {
		res := make([][]cards.Card, len(superset))
		for i, c := range superset {
			res[i] = []cards.Card{c}
		}
		return res
	}
	res := make([][]cards.Card, 0)
	for i, card := range superset {
		if i > len(superset)-n {
			break
		}
		others := superset[i+1:]
		combs := Combinations(others, n-1)
		for _, comb := range combs {
			set := append([]cards.Card{card}, comb...)
			res = append(res, set)
		}
	}
	return res
}

// allCombos is sized appropriately to fit every unique combo of 5 cards froma  52 card deck
// 2598960 is 52 choose 5
type allCombos = [2598960][5]uint8

// AllFiveCardHandIndices produces the sets of indices that represent every unique combination of 5 cards
// This is very specific to 52 cards in a deck and 5 cards in a hand, but that's okay for this
func AllFiveCardHandIndices() allCombos {
	ns := [52]uint8{}
	for i := uint8(0); i < 52; i++ {
		ns[i] = i
	}
	var res allCombos
	n := 0
	for i := 0; i < 48; i++ {
		for j := i + 1; j < 49; j++ {
			for k := j + 1; k < 50; k++ {
				for l := k + 1; l < 51; l++ {
					for m := l + 1; m < 52; m++ {
						res[n] = [5]uint8{ns[i], ns[j], ns[k], ns[l], ns[m]}
						n++
					}
				}
			}
		}
	}
	return res
}
