package score

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cszczepaniak/cribbage-scorer/testutils"
	"github.com/stretchr/testify/assert"
)

func TestScoreFifteens(t *testing.T) {
	tests := []struct {
		desc     string
		hand     []string
		cut      string
		expScore int
		expErr   error
	}{{
		desc:     `no fifteens`,
		hand:     []string{`9d`, `4h`, `7s`, `9s`},
		cut:      `jh`,
		expScore: 0,
		expErr:   nil,
	}, {
		desc:     `a hand`,
		hand:     []string{`10s`, `5h`, `7s`, `9s`},
		cut:      `jh`,
		expScore: 4,
		expErr:   nil,
	}, {
		desc:     `a hand`,
		hand:     []string{`7h`, `4s`, `4d`, `4c`},
		cut:      `4h`,
		expScore: 12,
		expErr:   nil,
	}, {
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `5h`,
		expScore: 2,
		expErr:   nil,
	}}
	for _, tc := range tests {
		hand, cut, err := testutils.MakeHandAndCut(tc.hand, tc.cut)
		require.NoError(t, err)
		score, err := scoreFifteens(hand, cut)
		if tc.expErr != nil {
			assert.EqualError(t, err, tc.expErr.Error())
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, tc.expScore, score)
	}
}
func TestScorePairs(t *testing.T) {
	tests := []struct {
		desc     string
		hand     []string
		cut      string
		expScore int
		expErr   error
	}{{
		desc:     `a hand`,
		hand:     []string{`6h`, `4s`, `4d`, `8c`},
		cut:      `7s`,
		expScore: 2,
		expErr:   nil,
	}, {
		desc:     `a hand`,
		hand:     []string{`6h`, `4s`, `4d`, `4c`},
		cut:      `7s`,
		expScore: 6,
		expErr:   nil,
	}, {
		desc:     `a hand`,
		hand:     []string{`7h`, `4s`, `4d`, `5c`},
		cut:      `7s`,
		expScore: 4,
		expErr:   nil,
	}, {
		desc:     `a hand`,
		hand:     []string{`7h`, `4s`, `4d`, `4c`},
		cut:      `7s`,
		expScore: 8,
		expErr:   nil,
	}, {
		desc:     `a hand`,
		hand:     []string{`7h`, `4s`, `4d`, `4c`},
		cut:      `4h`,
		expScore: 12,
		expErr:   nil,
	}, {
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `5h`,
		expScore: 0,
		expErr:   nil,
	}}
	for _, tc := range tests {
		hand, cut, err := testutils.MakeHandAndCut(tc.hand, tc.cut)
		require.NoError(t, err)
		score, err := scorePairs(hand, cut)
		if tc.expErr != nil {
			assert.EqualError(t, err, tc.expErr.Error())
		} else {
			assert.NoError(t, err)
		}
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
		expErr   error
	}{{
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `4s`},
		cut:      `5h`,
		isCrib:   false,
		expScore: 0,
		expErr:   nil,
	}, {
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `5s`,
		isCrib:   false,
		expScore: 4,
		expErr:   nil,
	}, {
		desc:     `a hand`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `5h`,
		isCrib:   false,
		expScore: 5,
		expErr:   nil,
	}, {
		desc:     `a crib`,
		hand:     []string{`ah`, `2h`, `3h`, `4s`},
		cut:      `5h`,
		isCrib:   true,
		expScore: 0,
		expErr:   nil,
	}, {
		desc:     `a crib`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `5s`,
		isCrib:   true,
		expScore: 0,
		expErr:   nil,
	}, {
		desc:     `a crib`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`},
		cut:      `5h`,
		isCrib:   true,
		expScore: 5,
		expErr:   nil,
	}}
	for _, tc := range tests {
		hand, cut, err := testutils.MakeHandAndCut(tc.hand, tc.cut)
		require.NoError(t, err)
		score, err := scoreFlush(hand, cut, tc.isCrib)
		if tc.expErr != nil {
			assert.EqualError(t, err, tc.expErr.Error())
		} else {
			assert.NoError(t, err)
		}
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
	}}
	for _, tc := range tests {
		hand, cut, err := testutils.MakeHandAndCut(tc.hand, tc.cut)
		require.NoError(t, err)
		score := scoreRuns(hand, cut)
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
		score := scoreNobs(hand, cut)
		assert.Equal(t, tc.expScore, score)
	}
}
