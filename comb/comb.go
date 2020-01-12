package comb

import (
	"github.com/cszczepaniak/cribbage-scorer/cards"
)

var cache map[int]int

func init() {
	cache = make(map[int]int)
}

func Factorial(n int) int {
	res, ok := cache[n]
	if ok {
		return res
	}
	if n <= 1 {
		res = 1
	} else {
		res = n * Factorial(n-1)
	}
	cache[n] = res
	return res
}

func Nchoosek(n, k int) int {
	return Factorial(n) / (Factorial(k) * Factorial(n-k))
}

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
