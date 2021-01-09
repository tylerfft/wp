package a06prototype

import (
	"fmt"
)

type MessageBoxStu struct {
	Deco string
}

func (r *MessageBoxStu) Init(Deco string) {
	r.Deco = Deco
}

func (r *MessageBoxStu) Use(Str string) {
	for i := 0; i < len(Str)+4; i++ {
		fmt.Print(r.Deco)
	}
	fmt.Println()
	fmt.Println(r.Deco + " " + Str + " " + r.Deco)
	for i := 0; i < len(Str)+4; i++ {
		fmt.Print(r.Deco)
	}
	fmt.Println()
}
func (r *MessageBoxStu) Clone() (out ProductIntf) {
	var MessageBox MessageBoxStu
	MessageBox = *r
	return &MessageBox
}
