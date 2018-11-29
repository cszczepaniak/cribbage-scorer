package main

import (
	"fmt"
	"os"
	"../pkg/cards"
)

func main() {
	hand := make([]cards.Card, 0, 5)

	clargs := os.Args[1:]
	for _, code := range clargs {
		hand = append(hand, cards.CardMap[code])
	}

	for _, card := range hand {
		cards.Print(card)
	}
	fmt.Println(findFifteens(hand))
}

func fact(n int) int {
	if n > 0 {
		return n * fact(n - 1)
	}
	return 1
}

func ncomb(n, m int) int {
	return fact(n) / (fact(m) * fact(n - m))
}

func comb(m int, set []cards.Card) [][]cards.Card {
	n := len(set)
	res := make([]cards.Card, m)
	last := m - 1
	total := ncomb(n, m)
	ret := make([][]cards.Card, 0, total)
    var rc func(int, int)
    rc = func(i, next int) {
        for j := next; j < n; j++ {
			res[i] = set[j]
            if i == last {
				newArr := make([]cards.Card, m)
				for i, val := range res {
					newArr[i] = val
				}
				ret = append(ret, newArr)
            } else {
                rc(i+1, j+1)
            }
        }
        return
    }
	rc(0, 0)
	
	return ret
}

func findNPairs(hand []cards.Card) int {
	n := 0

	allPairs := comb(2, hand)
	for _, val := range allPairs {
		if val[0].Rank == val[1].Rank {
			n++
		}
	}

	return n
}

func findFifteens(hand []cards.Card) int {
	n := 0

	for i := 2; i <= len(hand); i++ {
		for _, cards := range comb(i, hand) {
			sum := 0
			for _, card := range cards {
				sum += card.Value
			}

			if sum == 15 {
				n++
			}
		}
	}

	return n
}
