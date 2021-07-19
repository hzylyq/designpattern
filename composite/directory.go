package composite

type directory struct {
	name     string
	children []entry
}

func NewDirectory(name string) *directory {
	return &directory{
		name: name,
	}
}

func (d *directory) Name() string {
	return d.name
}

func (d *directory) Size() int {
	var size int
	for _, ent := range d.children {
		size += ent.Size()
	}
	return size
}

func (d *directory) PrintList() {

}

func (d *directory) Add(ent entry) error {
	d.children = append(d.children, ent)
	return nil
}
