package composite

import "log"

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

func (d *directory) PrintList(prefix string) {
	log.Println(prefix + "/" + d.Name())
	for _, ent := range d.children {
		ent.PrintList(prefix + "/" + d.Name())
	}
}

func (d *directory) Add(ent entry) {
	d.children = append(d.children, ent)
}
