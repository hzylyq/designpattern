package bridge

import "testing"

func TestBridge(t *testing.T) {
	d1 := NewCommonDisplay(NewStringDisplayImpl("Hello China."))
	d2 := NewCountDisplay(NewStringDisplayImpl("Hello World."))
	d3 := NewCountDisplay(NewStringDisplayImpl("Hello, Universe."))

	d1.display()
	d2.display()
	d3.display()
	d3.multiDisplay(5)
}
