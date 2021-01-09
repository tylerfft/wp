package a11composite

import (
	"fmt"
)

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

func (r *FileStu) PrintList(prefix string) {
	fmt.Println(prefix + "/" + r.Name)
}
