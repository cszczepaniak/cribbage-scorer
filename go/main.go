package main

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"runtime"

	"github.com/cszczepaniak/cribbage-scorer/cards"
	"github.com/cszczepaniak/cribbage-scorer/comb"
	"github.com/cszczepaniak/cribbage-scorer/score"
)

func main() {
	scores, err := scoreAll()
	if err != nil {
		log.Fatal(err)
	}

	err = writeResult(scores)
	if err != nil {
		log.Fatal(err)
	}
}

func scoreAll() (map[int]int, error) {
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
			return nil, err
		}
	}
	return scores, nil
}

func writeResult(scores map[int]int) error {
	bs, err := json.Marshal(scores)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(`result.json`, bs, fs.ModePerm)
}
