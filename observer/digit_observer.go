package observer

import (
	"fmt"
	"time"
)

type DigitObserver struct {
}

func (d *DigitObserver) Update(n *NumberGenerator) {
	fmt.Printf("DigitObserver:%d\n", n.number)
	time.Sleep(100)
}
