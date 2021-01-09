package a11composite

type EntryIntf interface {
	GetName() (Name string)
	GetSize() (Size int)
	PrintList(prefix string)
}

type EntryBaseStu struct {
	Entry EntryIntf
}

func (r *EntryBaseStu) Init(Entry EntryIntf) {
	r.Entry = Entry
}
func (r *EntryBaseStu) GetName() (Name string) {
	return r.Entry.GetName()
}
func (r *EntryBaseStu) GetSize() (Size int) {
	return r.Entry.GetSize()
}
