package strategy

type winningStrategy struct {
	won bool
}

func (w *winningStrategy) nextHand() {

}

func (w *winningStrategy) study(win bool) {
	w.won = win
}
