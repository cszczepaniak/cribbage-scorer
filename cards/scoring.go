package cards

import (
	"sort"

	"../comb"
)

func combinations(m int, set []Card) [][]Card {
	n := len(set)
	res := make([]Card, m)
	last := m - 1
	total := comb.Ncomb(n, m)
	ret := make([][]Card, 0, total)
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			res[i] = set[j]
			if i == last {
				newArr := make([]Card, m)
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

func ScoreHand(hand Hand) int {
	return ScorePairs(hand) + ScoreFifteens(hand) + ScoreRuns(hand)
}

func ScorePairs(hand Hand) int {
	score := 0

	allPairs := combinations(2, hand.AllCards())
	for _, val := range allPairs {
		if val[0].Rank == val[1].Rank {
			score += 2
		}
	}

	return score
}

func ScoreFifteens(hand Hand) int {
	score := 0

	for i := 2; i <= len(hand.AllCards()); i++ {
		for _, cards := range combinations(i, hand.AllCards()) {

			sum := 0
			for _, card := range cards {
				sum += card.Value
			}

			if sum == 15 {
				score += 2
			}
		}
	}

	return score
}

func ScoreRuns(hand Hand) int {
	score := 0
	minIdx := 2
	for i := 5; i > minIdx; i-- {
		for _, cards := range combinations(i, hand.AllCards()) {
			sort.Slice(cards, func(i, j int) bool { return cards[i].Rank < cards[j].Rank })
			if isRun(cards) {
				score += i
				minIdx = i
			}
		}
	}

	return score
}

func isRun(cards []Card) bool {
	for i := 0; i < len(cards)-1; i++ {
		if cards[i].Rank+1 != cards[i+1].Rank {
			return false
		}
	}
	return true
}
