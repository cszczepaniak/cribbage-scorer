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
