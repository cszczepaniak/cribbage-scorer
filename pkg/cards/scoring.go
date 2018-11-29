package cards

import(
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

func FindNPairs(hand []Card) int {
	n := 0

	allPairs := combinations(2, hand)
	for _, val := range allPairs {
		if val[0].Rank == val[1].Rank {
			n++
		}
	}

	return n
}

func FindFifteens(hand []Card) int {
	n := 0

	for i := 2; i <= len(hand); i++ {
		for _, cards := range combinations(i, hand) {
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
