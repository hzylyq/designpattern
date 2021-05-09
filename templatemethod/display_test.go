package templatemethod

import "testing"

func Test(t *testing.T) {
	d1 := charDisplay{
		ch: "H",
	}
	
	dA := d{
		AbstractDisplay: &d1,
	}
	dA.display()
	
	d2 := stringDisplay{
		width: 10,
		str:   "hello world",
	}
	dA = d{
		AbstractDisplay: &d2,
	}
	dA.display()
}
