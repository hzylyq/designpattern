package bridge

type Display struct {
	Method DisplayImpl
}

func NewDisplay(impl DisplayImpl) *Display {
	return &Display{
		Method: impl,
	}
}

func (d *Display) open() {
	d.Method.rawOpen()
}

func (d *Display) print() {
	d.Method.rawPrint()
}

func (d *Display) close() {
	d.Method.rawClose()
}

func (d *Display) display() {
	d.open()
	d.print()
	d.close()
}
