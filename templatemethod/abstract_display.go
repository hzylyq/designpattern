package templatemethod

type AbstractDisplay interface {
	open()
	print()
	close()
}

type d struct {
	AbstractDisplay AbstractDisplay
}

func (d *d) display() {
	d.AbstractDisplay.open()
	for i := 0; i < 5; i++ {
		d.AbstractDisplay.print()
	}
	d.AbstractDisplay.close()
}
