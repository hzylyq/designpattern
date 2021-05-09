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
		str: "hello world",
	}
	dA = d{
		AbstractDisplay: &d2,
	}
	dA.display()

	d2 = stringDisplay{
		str: "你好，世界。",
	}
	dA.AbstractDisplay = &d2
	dA.display()
}
