package base

import (
	"sync"
)

var SingleProRsv ProRsvStu
var SingleProRsvOnce sync.Once

func GetSingleProRsvStu() *ProRsvStu {
	SingleProRsvOnce.Do(SingleProRsv.Init)
	return &SingleProRsv
}

type ProRsvStu struct {
	ProRsv map[int]int
}

func (r *ProRsvStu) Init() {
	r.ProRsv = make(map[int]int)
}

func (r *ProRsvStu) GetRsv(Pro int) (Rsv int) {
	if Kind, ok := r.ProRsv[Pro]; ok {
		Rsv = Kind
	}
	return
}
func (r *ProRsvStu) Add(Pro, Rsv int) {
	r.ProRsv[Pro] = Rsv
}
