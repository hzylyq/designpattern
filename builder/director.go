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
	d.Builder.makeString("从早上至下午")
	d.Builder.makeItems([]string{
		"早上好",
		"下午好",
	})
	d.Builder.makeString("晚上")
	d.Builder.makeItems([]string{
		"晚上好",
		"晚安",
		"再见",
	})
	d.Builder.close()
}
