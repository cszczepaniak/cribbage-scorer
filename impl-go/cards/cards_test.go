package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func makeHandFromStrings(strs []string) ([]Card, error) {
	hand := make([]Card, len(strs))
	for i, s := range strs {
		c, err := FromString(s)
		if err != nil {
			return nil, err
		}
		hand[i] = c
	}
	return hand, nil
}

func TestSortByRankAscending(t *testing.T) {
	tests := []struct {
		hand []string
		exp  []string
	}{{
		hand: []string{`10h`, `8h`, `6h`, `4h`},
		exp:  []string{`4h`, `6h`, `8h`, `10h`},
	}, {
		hand: []string{`10h`, `8h`, `8s`, `4h`},
		exp:  []string{`4h`, `8h`, `8s`, `10h`},
	}, {
		hand: []string{`10h`, `8h`, `6h`},
		exp:  []string{`6h`, `8h`, `10h`},
	}, {
		hand: []string{`10h`, `6h`, `8h`, `ks`, `as`},
		exp:  []string{`as`, `6h`, `8h`, `10h`, `ks`},
	}}
	for _, tc := range tests {
		hand, err := makeHandFromStrings(tc.hand)
		require.NoError(t, err)
		exp, err := makeHandFromStrings(tc.exp)
		require.NoError(t, err)
		sorted := SortByRankAscending(hand)
		assert.Equal(t, exp, sorted)
	}
}

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
		expError: ErrInvalidCardString(`aaaa`),
	}, {
		desc:     `invalid string`,
		str:      `a`,
		exp:      Card{},
		expError: ErrInvalidCardString(`a`),
	}}
	for _, tc := range tests {
		c, err := FromString(tc.str)
		if tc.expError != nil {
			assert.EqualError(t, err, tc.expError.Error())
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, tc.exp, c)
	}
}

func BenchmarkNewDeck(b *testing.B) {
	_ = NewDeck()
}
