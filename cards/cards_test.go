package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCardFromString(t *testing.T) {
	tests := []struct {
		desc     string
		str      string
		exp      Card
		expError error
	}{{
		desc:     `ace of hearts`,
		str:      `ah`,
		exp:      Card{Suit: Hearts, Rank: 1, Name: `ah`},
		expError: nil,
	}, {
		desc:     `ace of hearts`,
		str:      `Ah`,
		exp:      Card{Suit: Hearts, Rank: 1, Name: `ah`},
		expError: nil,
	}, {
		desc:     `ten of spades`,
		str:      `10S`,
		exp:      Card{Suit: Spades, Rank: 10, Name: `10s`},
		expError: nil,
	}, {
		desc:     `nine of diamonds`,
		str:      `9d`,
		exp:      Card{Suit: Diamonds, Rank: 9, Name: `9d`},
		expError: nil,
	}, {
		desc:     `two of clubs`,
		str:      `2C`,
		exp:      Card{Suit: Clubs, Rank: 2, Name: `2c`},
		expError: nil,
	}, {
		desc:     `invalid rank`,
		str:      `!c`,
		exp:      Card{},
		expError: ErrInvalidRank,
	}, {
		desc:     `invalid rank`,
		str:      `11c`,
		exp:      Card{},
		expError: ErrInvalidRank,
	}, {
		desc:     `invalid rank`,
		str:      `1c`,
		exp:      Card{},
		expError: ErrInvalidRank,
	}, {
		desc:     `invalid suit`,
		str:      `2e`,
		exp:      Card{},
		expError: ErrInvalidSuit,
	}, {
		desc:     `invalid string`,
		str:      `aaaa`,
		exp:      Card{},
		expError: ErrInvalidCardString,
	}, {
		desc:     `invalid string`,
		str:      `a`,
		exp:      Card{},
		expError: ErrInvalidCardString,
	}}
	for _, tc := range tests {
		c, err := NewCardFromString(tc.str)
		if tc.expError != nil {
			assert.EqualError(t, err, tc.expError.Error())
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, tc.exp, c)
	}
}
