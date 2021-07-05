package decorator

import "testing"

func TestMain(t *testing.T) {
	b1 := NewDisplay("Hello, world.")
	b2 := NewSideBorder(*b1, '#')
	b3 := NewFullBorder(b2)

}
