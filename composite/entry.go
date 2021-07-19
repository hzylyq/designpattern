package composite

type entry interface {
	Name() string
	Size() int
	PrintList()
	Add(entry) error
}
