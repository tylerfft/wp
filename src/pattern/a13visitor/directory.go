package a13visitor

type DirectoryStu struct {
	Name      string
	EntryList []EntryIntf
}

func (r *DirectoryStu) Init(Name string) {
	r.Name = Name
}
func (r *DirectoryStu) AddEntry(Entry EntryIntf) {
	r.EntryList = append(r.EntryList, Entry)
}

func (r *DirectoryStu) GetName() (Name string) {
	return r.Name
}
func (r *DirectoryStu) GetSize() (Size int) {
	Size = 0
	for _, v := range r.EntryList {
		Size += v.GetSize()
	}
	return
}
func (r *DirectoryStu) ToString() (Str string) {
	return
}
func (r *DirectoryStu) Accept(Visitor VisitorIntf) {
	Visitor.VisitDirectory(r)
}
