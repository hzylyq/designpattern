package composite

type entry interface {
	Name() string
	Size() int
	PrintList(string)
	Add(entry)
}
