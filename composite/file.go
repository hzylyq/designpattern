package composite

import "log"

type file struct {
	name string
	size int
}

func NewFile(name string, size int) *file {
	return &file{
		name: name,
		size: size,
	}
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Size() int {
	return f.size
}

func (f *file) PrintList(prefix string) {
	log.Println(prefix + "/" + f.Name())
}

func (f *file) Add(entry) {}
