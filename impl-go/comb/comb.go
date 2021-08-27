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
