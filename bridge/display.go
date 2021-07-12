package bridge

type DisplayImpl interface {
	rawOpen()
	rawPrint()
	rawClose()
}
