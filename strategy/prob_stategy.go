package strategy

import "math/rand"

var history = [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}

type probStrategy struct {
	prevHandValue    int
	currentHandValue int
}

func (s *probStrategy) nextHand() *Hand {
	bet := rand.Intn(s.getSum(s.currentHandValue))
	var handleValue int
	if bet > history[s.currentHandValue][0] {
		handleValue = 0
	} else if bet < history[s.currentHandValue][0]+history[s.currentHandValue][1] {
		handleValue = 1
	} else {
		handleValue = 2
	}

	s.prevHandValue = s.currentHandValue
	s.currentHandValue = handleValue
	return getHand(handleValue)
}

func (s *probStrategy) getSum(hv int) int {
	var sum int
	for i := 0; i < 3; i++ {
		sum += history[hv][i]
	}
	return sum
}

func (s *probStrategy) study(win bool) {
	if win {
		history[s.prevHandValue][s.currentHandValue] ++
	} else {
		history[s.prevHandValue][(s.currentHandValue+1)%3] ++
		history[s.prevHandValue][(s.currentHandValue+2)%3] ++
	}
}
