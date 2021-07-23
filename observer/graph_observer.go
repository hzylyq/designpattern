package observer

import (
	"fmt"
	"time"
)

type GraphObserver struct {
}

func (g *GraphObserver) Update(n *NumberGenerator) {
	fmt.Print("GraphObserver:")
	count := n.number
	for i := 0; i < count; i++ {
		fmt.Print("*")
	}
	fmt.Println("")
	time.Sleep(100)
}
