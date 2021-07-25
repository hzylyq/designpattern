package state

import "testing"

func TestState(t *testing.T) {
	frame := NewSafeFrame()
	for hour := 0; hour < 24; hour++ {
		frame.setClock(hour)
	}
}
