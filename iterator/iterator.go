package iterator

type iterator interface {
	hasNext() bool
	Next() *Book
}
