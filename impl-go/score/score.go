package score

import (
	"errors"

	"github.com/cszczepaniak/cribbage-scorer/cards"
)

var (
	ErrInvalidHandSize = errors.New(`a hand must have exactly four cards in it`)
)

func ScoreHand(hand []cards.Card, cut cards.Card, isCrib bool) (int, error) {
	return scoreHand(hand, cut, isCrib), nil
}

func scoreHand(hand []cards.Card, cut cards.Card, isCrib bool) int {
	allCards := [5]cards.Card{hand[0], hand[1], hand[2], hand[3], cut}
	rankCounts := rankCounts{}
	values := [5]int{}
	for i, c := range allCards {
		rankCounts[c.Rank]++
		values[i] = c.Value()
	}
	return scoreFifteens(values) + scoreFlush(hand, cut, isCrib) + scoreNobs(hand, cut, rankCounts) + scorePairs(rankCounts) + scoreRuns(hand, rankCounts)
}

func scoreRuns(hand []cards.Card, rankCounts rankCounts) int {
	for _, c := range hand {
		if _, ok := rankCounts.get(c.Rank - 1); ok {
			// this is already part of a previously calculated run; skip calculation
			continue
		}
		// we're at the potential beginning of a run
		runLen := 1
		mult := rankCounts[c.Rank]
		nextUp := c.Rank + 1
		for ct, ok := rankCounts.get(nextUp); ok; ct, ok = rankCounts.get(nextUp) {
			mult *= ct
			runLen++
			nextUp++
		}
		if runLen >= 3 {
			return runLen * mult
		}
	}
	return 0
}

func scoreNobs(hand []cards.Card, cut cards.Card, rankCounts rankCounts) int {
	if cut.Rank == 11 {
		// cut is a jack; nothing to do
		return 0
	}
	if _, ok := rankCounts.get(11); !ok {
		// no jacks; nothing to do
		return 0
	}
	for _, c := range hand {
		if c.Rank == 11 && c.Suit == cut.Suit {
			return 1
		}
	}
	return 0
}

func scoreFlush(hand []cards.Card, cut cards.Card, isCrib bool) int {
	for _, c := range hand[1:] {
		if c.Suit != hand[0].Suit {
			return 0
		}
	}
	// ok, the hand is all matching now

	if cut.Suit == hand[0].Suit {
		// everything matches, doesn't matter if it's the crib or not, we have 5 points
		return 5
	}

	// now we know only the hand matches
	if !isCrib {
		return 4
	}
	return 0
}

func scorePairs(rankCounts rankCounts) int {
	score := 0
	for _, val := range rankCounts {
		switch val {
		case 2:
			score += 2
		case 3:
			score += 6
		case 4:
			score += 12
		}
	}
	return score
}

func scoreFifteens(values [5]int) int {
	if (values[0]|values[1]|values[2]|values[3]|values[4])&1 == 0 {
		return 0
	}
	sum := values[0] + values[1] + values[2] + values[3] + values[4]
	if sum < 15 || sum > 46 {
		return 0
	}
	return howManyFifteens(0, values[:]) * 2
}

func howManyFifteens(sum int, hand []int) int {
	if sum == 15 {
		return 1
	}
	if sum > 15 {
		return 0
	}

	var res int
	for i, c := range hand {
		res += howManyFifteens(sum+c, hand[i+1:])
	}
	return res
}

type rankCounts [15]int

func newRankCounts(cs []cards.Card) rankCounts {
	rc := rankCounts{}
	for _, c := range cs {
		rc[c.Rank]++
	}
	return rc
}

func (rc rankCounts) get(rank int) (int, bool) {
	return rc[rank], rc[rank] > 0
}
