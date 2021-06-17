package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cszczepaniak/cribbage-scorer/cards"
	"github.com/cszczepaniak/cribbage-scorer/score"
)

var (
	cutStr  = flag.String("cut", "", "the cut card")
	handStr = flag.String("hand", "", "a comma-separated list of card string representing the hand (e.g. ah,as,ad,ac)")
	isCrib  = flag.Bool("iscrib", false, "whether or not this is a crib")
)

func main() {
	start := time.Now()
	hand := make([]cards.Card, 0, 4)
	flag.Parse()
	parts := strings.Split(*handStr, `,`)
	if len(parts) != 4 {
		log.Fatalf(`Must have 4 cards, got %s`, *handStr)
	}
	for _, a := range parts {
		c, err := cards.NewCardFromString(strings.TrimSpace(a))
		if err != nil {
			log.Fatal(err)
		}
		hand = append(hand, c)
	}
	cut, err := cards.NewCardFromString(*cutStr)
	if err != nil {
		log.Fatal(err)
	}
	scorer := &score.Scorer{}
	s, err := scorer.ScoreHand(hand, cut, *isCrib)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Score is: %d\n", s)
	fmt.Printf("time elapsed: %s", time.Since(start))
}
