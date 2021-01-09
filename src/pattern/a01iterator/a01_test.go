package a01iterator

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	var BookShelf BookShelfStu
	BookShelf.AddBook(BookStu{Name: "shi"})
	BookShelf.AddBook(BookStu{Name: "new"})
	BookShelf.AddBook(BookStu{Name: "delete"})

	BookShelfIter := BookShelf.Iterator()
	exist := BookShelfIter.HasNext()
	for exist {
		tmp := BookShelfIter.Next()
		if v, ok := tmp.(BookStu); ok {
			fmt.Println(v.GetName())
		}
		exist = BookShelfIter.HasNext()

	}
}
