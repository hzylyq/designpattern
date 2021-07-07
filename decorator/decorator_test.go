package decorator

import "testing"

func TestDecorator(t *testing.T) {
	b1 := NewDisplay("Hello, world")
	b2 := NewSideBorder(b1, '#')
	b3 := NewFullBorder(b2)

	b1.show()
	b2.show()
	b3.show()

	b4 := NewSideBorder(NewFullBorder(
		NewFullBorder(
			NewSideBorder(
				NewFullBorder(
					NewDisplay("Hello, world."),
				), '*'),
		),
	), '/')
	b4.show()
}
