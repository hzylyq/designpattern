package observer

type Observer interface {
	Update(n *NumberGenerator)
}
