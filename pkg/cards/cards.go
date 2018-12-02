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

func init() {
	dat, _ := ioutil.ReadFile("cards.json")
	json.Unmarshal(dat, &AllCards)

	for _, card := range AllCards {
		CardMap[card.Code] = card
	}
}

func Print(card Card) {
	fmt.Println(card.Name, "of", card.Suit)
}
