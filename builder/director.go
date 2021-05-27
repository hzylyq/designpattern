package builder

type director struct {
	Builder builder
}

func NewDirector(b builder) *director {
	return &director{
		Builder: b,
	}
}

func (d *director) construct() {
	d.Builder.makeTitle("Greeting")
	d.Builder.makeString("")
}
