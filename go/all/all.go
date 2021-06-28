package all

import "github.com/cszczepaniak/cribbage-scorer/cards"

func EmitCombinations(n int, max int, handChan chan<- []cards.Card, errChan chan<- error) {
	state := make([]int, 0, n)
	for state[n-1] < max {
		if !allUnique(state) {
			continue
		}
		hand := make([]cards.Card, len(state))
		for i, s := range state {
			c, err := cards.FromIndex(s)
			if err != nil {
				errChan <- err
				return
			}
			hand[i] = c
		}
		handChan <- hand
		state = incState(state, max, 0)
	}
}

func allUnique(state []int) bool {
	unique := make(map[int]struct{})
	for _, s := range state {
		if _, ok := unique[s]; ok {
			return false
		}
		unique[s] = struct{}{}
	}
	return true
}

func incState(state []int, maxState int, startIdx int) []int {
	if startIdx+1 == len(state) {
		state[startIdx]++
		return state
	}
	state[startIdx]++
	if state[startIdx] > maxState {
		if startIdx+1 == len(state) {
			return state
		}
		state[startIdx] = 0
		incState(state, maxState, startIdx+1)
	}
	return state
}
