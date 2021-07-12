package strategy

type IStrategy interface {
	nextHand()
	study(win bool)
}
