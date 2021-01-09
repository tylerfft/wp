package a13visitor

import (
	"fmt"
)

type VisitorIntf interface {
	Reset(Dir string)
	ToString() (Str string)
	VisitFile(Entry *FileStu)
	VisitDirectory(Entry *DirectoryStu)
}

type VisitorListStu struct {
	Dir string
}

func (r *VisitorListStu) Reset(Dir string) {
	r.Dir = Dir
}

func (r *VisitorListStu) VisitFile(File *FileStu) {
	fmt.Println(r.Dir + "/" + File.GetName())

}
func (r *VisitorListStu) VisitDirectory(Directory *DirectoryStu) {

	fmt.Println(r.Dir + "/" + Directory.GetName())

	var tmp string = r.Dir
	r.Reset(r.Dir + "/" + Directory.GetName())
	for _, v := range Directory.EntryList {
		v.Accept(r)
	}
	r.Reset(tmp)
}

func (r *VisitorListStu) ToString() (Str string) {
	return r.Dir
}
