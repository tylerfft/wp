package a01iterator

type BookStu struct {
	Name string
}

func (r *BookStu) Init(Name string) {
	r.Name = Name
}

func (r *BookStu) GetName() (Name string) {
	return r.Name
}
