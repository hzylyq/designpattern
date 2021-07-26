package proxy

import (
	"fmt"
)

type printable interface {
	setPrinterName(name string)
	printerName() string
	print(str string)
}

type printer struct {
	name string
}

func (p *printer) setPrinterName(name string) {
	p.name = name
}

func (p *printer) printerName() string {
	return p.name
}

func (p *printer) print(str string) {
	fmt.Println("=== ", p.name, " ===")
	fmt.Println(str)
}
