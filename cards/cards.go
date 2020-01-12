package cards

import (
	"errors"
	"sort"
	"strings"
)

var (
	ErrInvalidCardString = errors.New(`invalid card string`)
	ErrInvalidSuit       = errors.New(`invalid suit`)
	ErrInvalidRank       = errors.New(`invalid rank`)
)

type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

type Card struct {
	Suit Suit   `json:"suit"`
	Rank int    `json:"rank"`
	Name string `json:"name"`
}

func NewCardFromString(s string) (Card, error) {
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
		return -1, ErrInvalidCardString
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
		return 0, ErrInvalidCardString
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
	return c.Name
}
