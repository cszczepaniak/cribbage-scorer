package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/cszczepaniak/cribbage-scorer/cards"
	"github.com/cszczepaniak/cribbage-scorer/comb"
	"github.com/cszczepaniak/cribbage-scorer/score"
)

var (
	cutStr  = flag.String("cut", "", "the cut card")
	handStr = flag.String("hand", "", "a comma-separated list of card string representing the hand (e.g. ah,as,ad,ac)")
	isCrib  = flag.Bool("iscrib", false, "whether or not this is a crib")
)

func main() {
	start := time.Now()
	deck := make([]cards.Card, 52)
	for i := 0; i < 52; i++ {
		c, err := cards.FromIndex(i)
		if err != nil {
			log.Fatal(err)
		}
		deck[i] = c
	}
	all := comb.Combinations(deck, 5)

	nWorkers := runtime.NumCPU()
	chunkSize := (len(all) + nWorkers - 1) / nWorkers
	scoreChan, doneChan, errChan := make(chan int), make(chan int), make(chan error)
	scorer := score.NewSerialScorer()

	for i := 0; i < len(all); i += chunkSize {
		end := i + chunkSize

		if end > len(all) {
			end = len(all)
		}
		fmt.Println(len(all[i:end]))
		go func(hands [][]cards.Card) {
			defer func() {
				doneChan <- 1
			}()
			for _, h := range hands {
				for j, c := range h {
					st := append([]cards.Card{}, h[:j]...)
					s, err := scorer.ScoreHand(append(st, h[j+1:]...), c, false)
					if err != nil {
						errChan <- err
						return
					}
					scoreChan <- s
				}
			}
		}(all[i:end])
	}

	done := 0
	scores := make(map[int]int, 30)
	for done < nWorkers {
		select {
		case d := <-doneChan:
			done += d
		case s := <-scoreChan:
			scores[s]++
		case err := <-errChan:
			log.Fatal(err)
		}
	}
	fmt.Println(scores)
	fmt.Printf("time elapsed: %s", time.Since(start))

	return
	start = time.Now()
	hand := make([]cards.Card, 0, 4)
	flag.Parse()
	parts := strings.Split(*handStr, `,`)
	if len(parts) != 4 {
		log.Fatalf(`Must have 4 cards, got %s`, *handStr)
	}
	for _, a := range parts {
		c, err := cards.FromString(strings.TrimSpace(a))
		if err != nil {
			log.Fatal(err)
		}
		hand = append(hand, c)
	}
	cut, err := cards.FromString(*cutStr)
	if err != nil {
		log.Fatal(err)
	}
	scorer1 := score.NewSerialScorer()
	s, err := scorer1.ScoreHand(hand, cut, *isCrib)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Score is: %d\n", s)
}
