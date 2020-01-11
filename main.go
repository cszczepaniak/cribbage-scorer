package main

import (
	"flag"
	"fmt"
)

func main() {
	hand := make([]cards.Card, 0, 5)

	cutPtr := flag.String("cut", "", "the cut card")
	flag.Parse()

	clargs := flag.Args()
	for _, code := range clargs {
		hand = append(hand, cards.CardMap[code])
	}

	myHand := cards.Hand{Cards: hand, Cut: cards.CardMap[*cutPtr], IsCrib: false}
	fmt.Println(myHand.AllCards())
	fmt.Println("Score:", cards.ScoreHand(myHand))
}
