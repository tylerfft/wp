package a11composite

import (
	"fmt"
)

type DirectoryStu struct {
	Name      string
	EntryList []EntryIntf
}

func (r *DirectoryStu) Init(Name string) {
	r.Name = Name

}
func (r *DirectoryStu) GetName() (Name string) {

	return r.Name
}
func (r *DirectoryStu) GetSize() (Size int) {
	Size = 0
	for _, v := range r.EntryList {
		Size += v.GetSize()
	}
	return Size
}
func (r *DirectoryStu) AddEntry(Entry EntryIntf) {
	r.EntryList = append(r.EntryList, Entry)
}

func (r *DirectoryStu) PrintList(prefix string) {
	fmt.Println(prefix + "/" + r.Name)
	for _, v := range r.EntryList {
		v.PrintList(prefix + "/" + r.Name)
	}
}
