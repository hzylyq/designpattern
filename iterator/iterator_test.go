package iterator

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	book1 := &Book{
		Name: "Around the world in 80 days",
	}
	book2 := &Book{
		Name: "Bible",
	}
	book3 := &Book{
		Name: "Cinderella",
	}
	book4 := &Book{Name: "Daddy-Long-Legs"}

	bookCollection := &bookCollection{
		BookList: []*Book{book1, book2, book3, book4},
	}

	iterator := bookCollection.createIterator()
	for iterator.hasNext() {
		book := iterator.Next()
		fmt.Printf("Book is %v\n", book)
	}
}
