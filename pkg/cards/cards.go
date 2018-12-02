package cards

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var AllCards []Card
var CardMap = make(map[string]Card)

type Card struct {
	Suit  string `json:"suit"`
	Value int    `json:"value"`
	Rank  int    `json:"rank"`
	Name  string `json:"name"`
	Code  string `json:"code"`
}

type Hand struct {
	Cards  []Card
	Cut    Card
	IsCrib bool
}

func (h Hand) AllCards() []Card {
	return append(h.Cards, h.Cut)
}

func init() {
	dat, _ := ioutil.ReadFile("cards.json")
	json.Unmarshal(dat, &AllCards)

	for _, card := range AllCards {
		CardMap[card.Code] = card
	}
}

func (c Card) Print() {
	fmt.Println(c.Name, "of", c.Suit)
}
