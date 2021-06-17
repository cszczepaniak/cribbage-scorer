package score

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cszczepaniak/cribbage-scorer/testutils"
	"github.com/stretchr/testify/assert"
)

func TestScoreHand(t *testing.T) {
	tests := []struct {
		desc     string
		hand     []string
		cut      string
		isCrib   bool
		expScore int
		expErr   error
	}{{
		desc:     `too many cards`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`, `5h`},
		cut:      `6h`,
		isCrib:   false,
		expScore: 0,
		expErr:   ErrInvalidHandSize,
	}, {
		desc:     `too few cards`,
		hand:     []string{`ah`, `2h`, `3h`},
		cut:      `6h`,
		isCrib:   false,
		expScore: 0,
		expErr:   ErrInvalidHandSize,
	}, {
		desc:     `perfect hand`,
		hand:     []string{`5c`, `5d`, `5h`, `js`},
		cut:      `5s`,
		isCrib:   false,
		expScore: 29,
		expErr:   nil,
	}, {
		desc:     `a really good hand`,
		hand:     []string{`4c`, `4d`, `5c`, `5d`},
		cut:      `6s`,
		isCrib:   false,
		expScore: 24,
		expErr:   nil,
	}, {
		desc:     `a really good crib`,
		hand:     []string{`4c`, `4d`, `5c`, `5d`},
		cut:      `6s`,
		isCrib:   true,
		expScore: 24,
		expErr:   nil,
	}, {
		desc:     `a really bad hand`,
		hand:     []string{`2c`, `4d`, `6c`, `8d`},
		cut:      `10d`,
		isCrib:   false,
		expScore: 0,
		expErr:   nil,
	}, {
		desc:     `a good hand`,
		hand:     []string{`10c`, `jc`, `jh`, `qc`},
		cut:      `5h`,
		isCrib:   false,
		expScore: 17,
		expErr:   nil,
	}, {
		desc:     `a hand with a flush`,
		hand:     []string{`6h`, `7h`, `8h`, `9h`},
		cut:      `10c`,
		isCrib:   false,
		expScore: 13,
		expErr:   nil,
	}, {
		desc:     `a crib with almost a flush`,
		hand:     []string{`6h`, `7h`, `8h`, `9h`},
		cut:      `10c`,
		isCrib:   true,
		expScore: 9,
		expErr:   nil,
	}, {
		desc:     `a crib with a flush`,
		hand:     []string{`6h`, `7h`, `8h`, `9h`},
		cut:      `10h`,
		isCrib:   true,
		expScore: 14,
		expErr:   nil,
	}, {
		desc:     `only nobs`,
		hand:     []string{`4h`, `6c`, `8s`, `jh`},
		cut:      `10h`,
		isCrib:   false,
		expScore: 1,
		expErr:   nil,
	}}
	for _, tc := range tests {
		hand, cut, err := testutils.MakeHandAndCut(tc.hand, tc.cut)
		require.NoError(t, err)
		scorer := &Scorer{}
		score, err := scorer.ScoreHand(hand, cut, tc.isCrib)
		if tc.expErr != nil {
			assert.EqualError(t, err, tc.expErr.Error())
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, tc.expScore, score)
	}
}

func TestScoreFifteens(t *testing.T) {
	tests := []struct {
		desc     string
		hand     []string
		cut      string
		expScore int
	}{{
		desc:     `no fifteens`,
		hand:     []string{`9d`, `4h`, `7s`, `9s`},
		cut:      `jh`,
		expScore: 0,
	}, {
		desc:     `a hand`,
		hand:     []string{`10s`, `5h`, `7s`, `9s`},
		cut:      `jh`,
		expScore: 4,
	}, {
		desc:     `a hand`,
		hand:     []string{`7h`, `4s`, `4d`, `4c`},
		cut:      `4h`,
		expScore: 12,
	}, {
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `5h`,
		expScore: 2,
	}}
	for _, tc := range tests {
		hand, cut, err := testutils.MakeHandAndCut(tc.hand, tc.cut)
		require.NoError(t, err)
		scorer := &Scorer{}
		score := scorer.scoreFifteens(hand, cut)
		assert.Equal(t, tc.expScore, score)
	}
}
func TestScorePairs(t *testing.T) {
	tests := []struct {
		desc     string
		hand     []string
		cut      string
		expScore int
	}{{
		desc:     `a hand`,
		hand:     []string{`6h`, `4s`, `4d`, `8c`},
		cut:      `7s`,
		expScore: 2,
	}, {
		desc:     `a hand`,
		hand:     []string{`6h`, `4s`, `4d`, `4c`},
		cut:      `7s`,
		expScore: 6,
	}, {
		desc:     `a hand`,
		hand:     []string{`7h`, `4s`, `4d`, `5c`},
		cut:      `7s`,
		expScore: 4,
	}, {
		desc:     `a hand`,
		hand:     []string{`7h`, `4s`, `4d`, `4c`},
		cut:      `7s`,
		expScore: 8,
	}, {
		desc:     `a hand`,
		hand:     []string{`7h`, `4s`, `4d`, `4c`},
		cut:      `4h`,
		expScore: 12,
	}, {
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `5h`,
		expScore: 0,
	}}
	for _, tc := range tests {
		hand, cut, err := testutils.MakeHandAndCut(tc.hand, tc.cut)
		require.NoError(t, err)
		scorer := &Scorer{}
		score := scorer.scorePairs(hand, cut)
		assert.Equal(t, tc.expScore, score)
	}
}
func TestScoreFlush(t *testing.T) {
	tests := []struct {
		desc     string
		hand     []string
		cut      string
		isCrib   bool
		expScore int
	}{{
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `4s`},
		cut:      `5h`,
		isCrib:   false,
		expScore: 0,
	}, {
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `5s`,
		isCrib:   false,
		expScore: 4,
	}, {
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `5h`,
		isCrib:   false,
		expScore: 5,
	}, {
		desc:     `a crib`,
		hand:     []string{`ah`, `2h`, `3h`, `4s`},
		cut:      `5h`,
		isCrib:   true,
		expScore: 0,
	}, {
		desc:     `a crib`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `5s`,
		isCrib:   true,
		expScore: 0,
	}, {
		desc:     `a crib`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `5h`,
		isCrib:   true,
		expScore: 5,
	}}
	for _, tc := range tests {
		hand, cut, err := testutils.MakeHandAndCut(tc.hand, tc.cut)
		require.NoError(t, err)
		scorer := &Scorer{}
		score := scorer.scoreFlush(hand, cut, tc.isCrib)
		assert.Equal(t, tc.expScore, score)
	}
}
func TestScoreRuns(t *testing.T) {
	tests := []struct {
		desc     string
		hand     []string
		cut      string
		expScore int
	}{{
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `5h`,
		expScore: 5,
	}, {
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `9h`,
		expScore: 4,
	}, {
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `8h`},
		cut:      `9h`,
		expScore: 3,
	}, {
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `4s`,
		expScore: 8,
	}, {
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `3s`},
		cut:      `10h`,
		expScore: 6,
	}, {
		desc:     `a hand`,
		hand:     []string{`10c`, `jc`, `jh`, `qc`},
		cut:      `ah`,
		expScore: 6,
	}}
	for _, tc := range tests {
		hand, cut, err := testutils.MakeHandAndCut(tc.hand, tc.cut)
		require.NoError(t, err)
		scorer := &Scorer{}
		score := scorer.scoreRuns(hand, cut)
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
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `jh`},
		cut:      `4h`,
		expScore: 1,
	}, {
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `jh`},
		cut:      `4c`,
		expScore: 0,
	}, {
		desc:     `a hand`,
		hand:     []string{`jc`, `jd`, `jh`, `js`},
		cut:      `4c`,
		expScore: 1,
	}}
	for _, tc := range tests {
		hand, cut, err := testutils.MakeHandAndCut(tc.hand, tc.cut)
		require.NoError(t, err)
		scorer := &Scorer{}
		score := scorer.scoreNobs(hand, cut)
		assert.Equal(t, tc.expScore, score)
	}
}
