package a06prototype

import (
	"fmt"
)

type UnderLineStu struct {
	Unch string
}

func (r *UnderLineStu) Init(Unch string) {
	r.Unch = Unch
}

func (r *UnderLineStu) Use(Str string) {

	fmt.Println("\"" + Str + "\"")
	fmt.Print(" ")
	for i := 0; i < len(Str); i++ {
		fmt.Print(r.Unch)
	}
	fmt.Println(" ")

}
func (r *UnderLineStu) Clone() (out ProductIntf) {
	var UnderLine UnderLineStu
	UnderLine = *r
	return &UnderLine
}
