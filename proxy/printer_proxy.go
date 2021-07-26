package proxy

type printerProxy struct {
	name string
	real *printer
}

func (p *printerProxy) PrintName() string {
	return p.name
}

func (p *printerProxy) Print(str string) {
	p.realize()
	p.real.print(str)
}

func (p *printerProxy) realize() {
	if p.real == nil {
		p.real = NewPrinter(p.name)
	}
}
