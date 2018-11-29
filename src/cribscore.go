package main

import (
	"fmt"
	"os"
	"../pkg/cards"
)

func main() {
	hand := make([]cards.Card, 0, 5)

	clargs := os.Args[1:]
	for _, code := range clargs {
		hand = append(hand, cards.CardMap[code])
	}

	for _, card := range hand {
		cards.Print(card)
	}
	fmt.Println(cards.FindNPairs(hand))
}
