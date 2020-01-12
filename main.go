package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/cszczepaniak/cribbage-scorer/cards"
	"github.com/cszczepaniak/cribbage-scorer/score"
)

var (
	cutStr = flag.String("cut", "", "the cut card")
	isCrib = flag.Bool("iscrib", false, "whether or not this is a crib")
)

func main() {
	start := time.Now()
	hand := make([]cards.Card, 0, 4)
	flag.Parse()
	args := flag.Args()
	for _, a := range args {
		c, err := cards.NewCardFromString(a)
		if err != nil {
			log.Fatal(err)
		}
		hand = append(hand, c)
	}
	cut, err := cards.NewCardFromString(*cutStr)
	if err != nil {
		log.Fatal(err)
	}
	s, err := score.ScoreHand(hand, cut, *isCrib)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Score is: %d\n", s)
	fmt.Printf("time elapsed: %s", time.Since(start))
}
