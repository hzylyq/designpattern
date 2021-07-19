package visit

type entry interface {
	Name() string
	Size() int
	Accept(Visitor)
}
