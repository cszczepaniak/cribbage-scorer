package cards

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var AllCards []Card
var CardMap map[string]Card

type Card struct {
	Suit  string `json:"suit"`
	Value int    `json:"value"`
	Rank  string `json:"rank"`
	Code  string `json:"code"`
}

func init() {
	dat, _ := ioutil.ReadFile("../cards.json")
	json.Unmarshal(dat, &AllCards)

	CardMap := make(map[string]Card)

	for _, card := range AllCards {
		CardMap[card.Code] = card
	}
}

func Print(card Card) {
	fmt.Println(card.Rank, "of", card.Suit)
}