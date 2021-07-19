package visit

type Directory struct {
	name     string
	children []entry
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name: name,
	}
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Size() int {
	var size int
	for _, child := range d.children {
		size += child.Size()
	}
	return size
}

func (d *Directory) Add(e entry) {
	d.children = append(d.children, e)
}

func (d *Directory) Accept(v Visitor) {
	v.Visit(d)
}
