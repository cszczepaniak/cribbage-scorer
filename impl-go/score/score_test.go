package score

import (
	"testing"

	"github.com/cszczepaniak/cribbage-scorer/testutils"
	"github.com/stretchr/testify/assert"
)

func BenchmarkScoreHand(b *testing.B) {
	tests := []struct {
		desc     string
		hand     []string
		cut      string
		isCrib   bool
		expScore int
	}{{
		desc:     `perfect hand`,
		hand:     []string{`5c`, `5d`, `5h`, `js`},
		cut:      `5s`,
		isCrib:   false,
		expScore: 29,
	}, {
		desc:     `a really good hand`,
		hand:     []string{`4c`, `4d`, `5c`, `5d`},
		cut:      `6s`,
		isCrib:   false,
		expScore: 24,
	}, {
		desc:     `a really good crib`,
		hand:     []string{`4c`, `4d`, `5c`, `5d`},
		cut:      `6s`,
		isCrib:   true,
		expScore: 24,
	}, {
		desc:     `a really bad hand`,
		hand:     []string{`2c`, `4d`, `6c`, `8d`},
		cut:      `10d`,
		isCrib:   false,
		expScore: 0,
	}, {
		desc:     `a good hand`,
		hand:     []string{`10c`, `jc`, `jh`, `qc`},
		cut:      `5h`,
		isCrib:   false,
		expScore: 17,
	}, {
		desc:     `a hand with a flush`,
		hand:     []string{`6h`, `7h`, `8h`, `9h`},
		cut:      `10c`,
		isCrib:   false,
		expScore: 13,
	}, {
		desc:     `a crib with almost a flush`,
		hand:     []string{`6h`, `7h`, `8h`, `9h`},
		cut:      `10c`,
		isCrib:   true,
		expScore: 9,
	}, {
		desc:     `a crib with a flush`,
		hand:     []string{`6h`, `7h`, `8h`, `9h`},
		cut:      `10h`,
		isCrib:   true,
		expScore: 14,
	}, {
		desc:     `only nobs`,
		hand:     []string{`4h`, `6c`, `8s`, `jh`},
		cut:      `10h`,
		isCrib:   false,
		expScore: 1,
	}}
	for _, tc := range tests {
		tc := tc
		hand, cut := testutils.MakeHandAndCut(b, tc.hand, tc.cut)

		b.Run(tc.desc, func(b *testing.B) {
			score := ScoreHand(hand, cut, tc.isCrib)
			assert.Equal(b, tc.expScore, score)
		})
	}
}

func TestScoreFifteens(t *testing.T) {
	tests := []struct {
		desc     string
		values   [5]int
		expScore int
	}{{
		desc:     `no fifteens`,
		values:   [5]int{9, 4, 7, 9, 10},
		expScore: 0,
	}, {
		desc:     `a hand`,
		values:   [5]int{10, 5, 7, 9, 10},
		expScore: 4,
	}, {
		desc:     `a hand`,
		values:   [5]int{7, 4, 4, 4, 4},
		expScore: 12,
	}, {
		desc:     `a hand`,
		values:   [5]int{1, 2, 3, 4, 5},
		expScore: 2,
	}}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, tc.expScore, scoreFifteens(tc.values))
		})
	}
}
func TestScorePairs(t *testing.T) {
	tests := []struct {
		hand     []string
		expScore int
	}{{
		hand:     []string{`6h`, `4s`, `4d`, `8c`, `7s`},
		expScore: 2,
	}, {
		hand:     []string{`6h`, `4s`, `4d`, `4c`, `7s`},
		expScore: 6,
	}, {
		hand:     []string{`7h`, `4s`, `4d`, `5c`, `7s`},
		expScore: 4,
	}, {
		hand:     []string{`7h`, `4s`, `4d`, `4c`, `7s`},
		expScore: 8,
	}, {
		hand:     []string{`7h`, `4s`, `4d`, `4c`, `4h`},
		expScore: 12,
	}, {
		hand:     []string{`ah`, `2h`, `3h`, `4h`, `5h`},
		expScore: 0,
	}}
	for _, tc := range tests {
		cards := testutils.ParseCards(t, tc.hand)
		score := scorePairs(newRankCounts(cards))
		assert.Equal(t, tc.expScore, score)
	}
}
func TestScoreFlush(t *testing.T) {
	tests := []struct {
		hand         []string
		cut          string
		expHandScore int
		expCribScore int
	}{{
		hand:         []string{`ah`, `2h`, `3h`, `4s`},
		cut:          `5h`,
		expHandScore: 0,
		expCribScore: 0,
	}, {
		hand:         []string{`ah`, `2h`, `3h`, `4h`},
		cut:          `5s`,
		expHandScore: 4,
		expCribScore: 0,
	}, {
		hand:         []string{`ah`, `2h`, `3h`, `4h`},
		cut:          `5h`,
		expHandScore: 5,
		expCribScore: 5,
	}, {
		hand:         []string{`ah`, `2h`, `3h`, `4s`},
		cut:          `5h`,
		expHandScore: 0,
		expCribScore: 0,
	}}
	for _, tc := range tests {
		hand, cut := testutils.MakeHandAndCut(t, tc.hand, tc.cut)
		assert.Equal(t, tc.expHandScore, scoreFlush(hand, cut, false))
		assert.Equal(t, tc.expCribScore, scoreFlush(hand, cut, true))
	}
}
func TestScoreRuns(t *testing.T) {
	tests := []struct {
		hand     []string
		expScore int
	}{{
		hand:     []string{`ah`, `2h`, `3h`, `4h`, `5h`},
		expScore: 5,
	}, {
		hand:     []string{`ah`, `2h`, `3h`, `4h`, `9h`},
		expScore: 4,
	}, {
		hand:     []string{`ah`, `2h`, `3h`, `8h`, `9h`},
		expScore: 3,
	}, {
		hand:     []string{`ah`, `2h`, `3h`, `4h`, `4s`},
		expScore: 8,
	}, {
		hand:     []string{`ah`, `2h`, `3h`, `3s`, `10h`},
		expScore: 6,
	}, {
		hand:     []string{`10c`, `jc`, `jh`, `qc`, `ah`},
		expScore: 6,
	}}
	for _, tc := range tests {
		hand := testutils.ParseCards(t, tc.hand)
		score := scoreRuns(hand, newRankCounts(hand))
		assert.Equal(t, tc.expScore, score)
	}
}
func TestScoreNobs(t *testing.T) {
	tests := []struct {
		desc     string
		hand     []string
		cut      string
		expScore int
	}{{
		hand:     []string{`ah`, `2h`, `3h`, `jh`, `4h`},
		expScore: 1,
	}, {
		hand:     []string{`ah`, `2h`, `3h`, `jh`, `4c`},
		expScore: 0,
	}, {
		hand:     []string{`jc`, `jd`, `jh`, `js`, `4c`},
		expScore: 1,
	}}
	for _, tc := range tests {
		hand := testutils.ParseCards(t, tc.hand)
		score := scoreNobs(hand[0:4], hand[4], newRankCounts(hand))
		assert.Equal(t, tc.expScore, score)
	}
}
