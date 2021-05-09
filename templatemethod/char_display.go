package templatemethod

import "fmt"

type charDisplay struct {
	ch string
}

func (c *charDisplay) open() {
	fmt.Print("<<")
}

func (c *charDisplay) print() {
	fmt.Print(c.ch)
}

func (c *charDisplay) close() {
	fmt.Print(">>")
}
