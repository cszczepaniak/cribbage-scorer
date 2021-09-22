package main

import (
	"fmt"
	"log"
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
	deck := cards.NewDeck()
	allIndices := comb.AllFiveCardHandIndices()

	nWorkers := runtime.NumCPU()
	chunkSize := (len(allIndices) + nWorkers - 1) / nWorkers

	var scores [30]int32
	var wg sync.WaitGroup
	wg.Add(nWorkers)

	for i := 0; i < len(allIndices); i += chunkSize {
		end := i + chunkSize
		if end > len(allIndices) {
			end = len(allIndices)
		}

		go func(start, end int) {
			defer func() {
				wg.Done()
			}()
			for i := start; i < end; i++ {
				idxs := allIndices[i]
				// for each set of 5, there are 5 hands we can build
				hs := [5][5]cards.Card{
					{deck[idxs[0]], deck[idxs[1]], deck[idxs[2]], deck[idxs[3]], deck[idxs[4]]},
					{deck[idxs[4]], deck[idxs[0]], deck[idxs[1]], deck[idxs[2]], deck[idxs[3]]},
					{deck[idxs[3]], deck[idxs[4]], deck[idxs[0]], deck[idxs[1]], deck[idxs[2]]},
					{deck[idxs[2]], deck[idxs[3]], deck[idxs[4]], deck[idxs[0]], deck[idxs[1]]},
					{deck[idxs[1]], deck[idxs[2]], deck[idxs[3]], deck[idxs[4]], deck[idxs[0]]},
				}
				for _, h := range hs {
					s := score.ScoreHand(h[0:4], h[4], false)
					atomic.AddInt32(&scores[s], 1)
				}
			}
		}(i, end)
	}

	wg.Wait()
	return scores, nil
}
