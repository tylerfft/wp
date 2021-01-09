package a14chain

import (
	"strconv"
)

type TroubleStu struct {
	Num int
}

func (r *TroubleStu) Init(Num int) {
	r.Num = Num
}

func (r *TroubleStu) Get() (Num int) {
	return r.Num
}
func (r *TroubleStu) ToString() (Str string) {
	return "[ Trouble   " + strconv.Itoa(r.Num) + "]"
}
