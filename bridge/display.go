package bridge

type Display interface {
	open()
	print()
	close()
	display()
}

type DisplayImpl interface {
	rawOpen()
	rawPrint()
	rawClose()
}


