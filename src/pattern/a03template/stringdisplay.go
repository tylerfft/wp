package a03template

import (
	"fmt"
)

type StringDisplayStu struct {
	str string
}

func (r *StringDisplayStu) Init(str string) {
	r.str = str
}

func (r *StringDisplayStu) Open() {
	r.PrintLine()
}
func (r *StringDisplayStu) Print() {
	fmt.Println("|" + r.str + "|")
}

func (r *StringDisplayStu) Close() {
	r.PrintLine()
}

func (r *StringDisplayStu) PrintLine() {
	fmt.Print("+")
	for i := 0; i < len(r.str); i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
