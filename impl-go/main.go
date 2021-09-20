package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/cszczepaniak/cribbage-scorer/cards"
	"github.com/cszczepaniak/cribbage-scorer/comb"
	"github.com/cszczepaniak/cribbage-scorer/score"
)

func main() {
	scores, err := scoreAll()
	if err != nil {
		log.Fatal(err)
	}

	for s, ct := range scores {
		fmt.Printf("%02d: %d\n", s, ct)
	}
}

func scoreAll() ([30]int32, error) {
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
	scorer := score.NewSerialScorer()

	var scores [30]int32
	var wg sync.WaitGroup
	wg.Add(nWorkers)

	for i := 0; i < len(all); i += chunkSize {
		end := i + chunkSize

		if end > len(all) {
			end = len(all)
		}
		go func(hands [][]cards.Card) {
			defer func() {
				wg.Done()
			}()
			for _, h := range hands {
				for j, c := range h {
					st := append([]cards.Card{}, h[:j]...)
					s, err := scorer.ScoreHand(append(st, h[j+1:]...), c, false)
					if err != nil {
						fmt.Fprintf(os.Stderr, "%v\n", err)
						return
					}
					atomic.AddInt32(&scores[s], 1)
				}
			}
		}(all[i:end])
	}

	wg.Wait()
	return scores, nil
}
