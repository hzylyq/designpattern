package iterator

type bookIterator struct {
	Index    int
	BookList []*Book
}

func (u *bookIterator) hasNext() bool {
	if u.Index < len(u.BookList) {
		return true
	}
	return false
}

func (u *bookIterator) Next() *Book {
	if u.hasNext() {
		book := u.BookList[u.Index]
		u.Index++
		return book
	}
	return nil
}
