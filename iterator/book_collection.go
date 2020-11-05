package iterator

type bookCollection struct {
	BookList []*Book
}

func (coll *bookCollection) createIterator() iterator {
	return &bookIterator{
		BookList: coll.BookList,
	}
}
