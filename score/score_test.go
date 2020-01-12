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
	}, {
		desc:     `hand too small`,
		hand:     []string{`ah`, `2h`, `3h`},
		cut:      `5h`,
		expScore: 0,
		expErr:   ErrInvalidHandSize,
	}, {
		desc:     `hand too big`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`, `6h`},
		cut:      `5h`,
		expScore: 0,
		expErr:   ErrInvalidHandSize,
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
	}, {
		desc:     `hand too small`,
		hand:     []string{`ah`, `2h`, `3h`},
		cut:      `5h`,
		expScore: 0,
		expErr:   ErrInvalidHandSize,
	}, {
		desc:     `hand too big`,
		hand:     []string{`ah`, `2h`, `3h`, `4h`, `6h`},
		cut:      `5h`,
		expScore: 0,
		expErr:   ErrInvalidHandSize,
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
