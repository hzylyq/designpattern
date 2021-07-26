package proxy

import (
	"fmt"
	"time"
)

type printable interface {
	setPrinterName(name string)
	printerName() string
	print(str string)
}

type printer struct {
	name string
}

func NewPrinter(name string) *printer {
	p := &printer{
		name: name,
	}
	p.heavyJob("正在生成printer的实例")
	return p
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

func (p *printer) heavyJob(msg string) {
	fmt.Print(msg)
	for i := 0; i < 5; i++ {
		time.Sleep(1000)
		fmt.Print(".")
	}
	fmt.Println("结束.")
}
