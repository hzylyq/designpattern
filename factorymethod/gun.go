package factorymethod

type gun struct {
	name  string
	power int
}

func (g *gun) Name() string {
	return g.name
}

func (g *gun) SetName(name string) {
	g.name = name
}

func (g *gun) Power() int {
	return g.power
}

func (g *gun) SetPower(power int) {
	g.power = power
}
