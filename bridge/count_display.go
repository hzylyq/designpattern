package bridge

type CountDisplay struct {
	CommonDisplay
}

func NewCountDisplay(impl DisplayImpl) *CountDisplay {
	return &CountDisplay{
		CommonDisplay{
			Method: impl,
		},
	}
}

func (d *CountDisplay) multiDisplay(times int) {
	d.open()
	for i := 0; i < times; i++ {
		d.print()
	}
	d.close()
}
