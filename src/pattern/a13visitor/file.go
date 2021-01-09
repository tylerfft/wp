package a13visitor

type FileStu struct {
	Name string
	Size int
}

func (r *FileStu) Init(Name string, Size int) {
	r.Name = Name
	r.Size = Size
}

func (r *FileStu) GetName() (Name string) {
	return r.Name
}
func (r *FileStu) GetSize() (Size int) {
	return r.Size
}
func (r *FileStu) ToString() (Str string) {
	return
}
func (r *FileStu) Accept(Visitor VisitorIntf) {
	Visitor.VisitFile(r)
}
