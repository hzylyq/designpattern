package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	p := NewPrinterProxy("Alice")
	fmt.Println(p.PrintName())
	p.SetPrintName("Bob")
	fmt.Println(p.PrintName())
	p.Print("Hello, world.")
}
