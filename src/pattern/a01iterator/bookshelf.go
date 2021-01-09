package a01iterator

type BookShelfStu struct {
	Books []BookStu
}

func (r *BookShelfStu) AddBook(Books BookStu) {
	r.Books = append(r.Books, Books)
}

func (r *BookShelfStu) GetLength() (Cnt int) {
	return len(r.Books)
}

func (r *BookShelfStu) Iterator() (Iter IterIntf) {
	var BookShelfIter BookShelfIterStu
	BookShelfIter.Init(r)
	Iter = &BookShelfIter
	return
}
func (r *BookShelfStu) GetByIdx(Idx int) (Book BookStu) {
	if Idx < len(r.Books) {
		Book = r.Books[Idx]
	}
	return
}
