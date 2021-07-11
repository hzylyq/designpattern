package bridge

type CommonDisplay struct {
	Method DisplayImpl
}

func NewCommonDisplay(impl DisplayImpl) *CommonDisplay {
	return &CommonDisplay{
		Method: impl,
	}
}

func (d *CommonDisplay) open() {
	d.Method.rawOpen()
}

func (d *CommonDisplay) print() {
	d.Method.rawPrint()
}

func (d *CommonDisplay) close() {
	d.Method.rawClose()
}

func (d *CommonDisplay) display() {
	d.open()
	d.print()
	d.close()
}
