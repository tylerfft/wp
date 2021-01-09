package tool

import (
	"fmt"
	"strings"
)

func NewMyflagStu(Data string) *MyflagStu {
	var Myflag MyflagStu
	Myflag.Init(Data)
	return &Myflag
}

type MyflagStu struct {
	Data string
	Args []string
}

func (r *MyflagStu) Init(Data string) {
	r.Data = Data
	r.Split()
}
func (r *MyflagStu) Split() {
	r.Args = strings.Split(r.Data, " ")
}
func (r *MyflagStu) Print() {
	for _, v := range r.Args {
		fmt.Println(v, ",")
	}
}
