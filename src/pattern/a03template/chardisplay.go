package a03template

import (
	"fmt"
)

type CharDisplayStu struct {
	ch string
}

func (r *CharDisplayStu) Init(ch string) {
	r.ch = ch
}

func (r *CharDisplayStu) Open() {
	fmt.Print("<<")
}
func (r *CharDisplayStu) Print() {
	fmt.Print(r.ch)
}

func (r *CharDisplayStu) Close() {
	fmt.Println(">>")
}
