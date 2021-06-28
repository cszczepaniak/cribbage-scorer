package cards

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	ErrInvalidSuit = errors.New(`invalid suit`)
	ErrInvalidRank = errors.New(`invalid rank`)
)

func ErrInvalidCardString(got string) error {
	return fmt.Errorf(`invalid card string: %s`, got)
}

type Suit int

func (s Suit) String() string {
	switch s {
	case Clubs:
		return `c`
	case Diamonds:
		return `d`
	case Hearts:
		return `h`
	case Spades:
		return `s`
	}
	return `wtf`
}

const (
	Clubs    Suit = 0
	Diamonds Suit = 1
	Hearts   Suit = 2
	Spades   Suit = 3
)

type Card struct {
	Suit Suit   `json:"suit"`
	Rank int    `json:"rank"`
	Name string `json:"name"`
}

func FromIndex(idx int) (Card, error) {
	if idx < 0 || idx > 51 {
		return Card{}, errors.New(`index must be between 0 and 51 inclusive`)
	}
	return Card{
		Suit: Suit(idx % 4),
		Rank: (idx / 4) + 1,
	}, nil
}

func FromString(s string) (Card, error) {
	rank, err := rankFromString(s)
	if err != nil {
		return Card{}, err
	}
	suit, err := suitFromString(s)
	if err != nil {
		return Card{}, err
	}
	return Card{
		Suit: suit,
		Rank: rank,
		Name: strings.ToLower(s),
	}, nil
}

func (c Card) Value() int {
	if c.Rank < 11 {
		return c.Rank
	}
	return 10
}

func SortByRankAscending(set []Card) []Card {
	setCopy := make([]Card, len(set))
	copy(setCopy, set)
	sort.Slice(setCopy, func(i, j int) bool {
		return setCopy[i].Rank < setCopy[j].Rank
	})
	return setCopy
}

func suitFromString(s string) (Suit, error) {
	s = strings.ToLower(s)
	runes := []rune(s)
	if len(runes) > 3 || len(runes) < 2 {
		return -1, ErrInvalidCardString(s)
	}
	switch runes[len(runes)-1] {
	case 'c':
		return Clubs, nil
	case 'd':
		return Diamonds, nil
	case 'h':
		return Hearts, nil
	case 's':
		return Spades, nil
	default:
		return -1, ErrInvalidSuit
	}
}

func rankFromString(s string) (int, error) {
	s = strings.ToLower(s)
	runes := []rune(s)
	if len(runes) > 3 || len(runes) < 2 {
		return 0, ErrInvalidCardString(s)
	}
	if len(runes) == 3 {
		if runes[0] == '1' && runes[1] == '0' {
			return 10, nil
		}
		return 0, ErrInvalidRank
	}
	switch runes[0] {
	case 'a':
		return 1, nil
	case 'j':
		return 11, nil
	case 'q':
		return 12, nil
	case 'k':
		return 13, nil
	default:
		if runes[0] < '2' || runes[0] > '9' {
			return 0, ErrInvalidRank
		}
		return int(runes[0] - '0'), nil
	}
}

func (c Card) String() string {
	if c.Name != `` {
		return c.Name
	}
	n := ``
	// if c.Rank >= 2 && c.Rank <= 10 {
	// 	n += strconv.Itoa(c.Rank)
	// }
	switch c.Rank {
	case 1:
		n += `a`
	case 11:
		n += `j`
	case 12:
		n += `q`
	case 13:
		n += `k`
	default:
		n += strconv.Itoa(c.Rank)
	}
	return n + c.Suit.String()
}
