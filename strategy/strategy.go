package strategy

type IStrategy interface {
	nextHand() *Hand
	study(win bool)
}
