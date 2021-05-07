package templatemethod

type AbstractDisplay interface {
	open()
	print()
	close()
	display()
}

type display struct {
	AbstractDisplay AbstractDisplay
}

func (d *display) display() {
	for i := 0; i < 5; i++ {
		d.AbstractDisplay.print()
	}
}
