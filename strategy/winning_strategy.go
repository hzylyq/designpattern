package strategy

import "math/rand"

type winningStrategy struct {
	won      bool
	prevHand *Hand
}

func (w *winningStrategy) nextHand() *Hand {
	if !w.won {
		w.prevHand = getHand(rand.Intn(3))
	}
	return w.prevHand
}

func (w *winningStrategy) study(win bool) {
	w.won = win
}
