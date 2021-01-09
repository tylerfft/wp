package base

import (
	"sync"
)

func NewEntry(Kind string) (Entry EntryIntf) {
	CreateFunc := GetMethodFactory()
	Entry = CreateFunc.Create(Kind)
	return
}

var SingleMethodFactory MethodFactoryStu
var SingleMethodFactoryOnce sync.Once

func GetMethodFactory() *MethodFactoryStu {
	SingleMethodFactoryOnce.Do(SingleMethodFactory.Init)
	return &SingleMethodFactory
}

type EntryCreateFunc func() (Entry EntryIntf)

type MethodFactoryStu struct {
	Factorys map[string]EntryCreateFunc
}

func (r *MethodFactoryStu) Create(Kind string) (Entry EntryIntf) {
	if CreateFunc, ok := r.Factorys[Kind]; ok {
		Entry = CreateFunc()
	} else {
		Entry = nil
	}
	return
}

func (r *MethodFactoryStu) Add(Kind string, CreateFunc EntryCreateFunc) {
	r.Factorys[Kind] = CreateFunc
}

func (r *MethodFactoryStu) Init() {
	r.Factorys = make(map[string]EntryCreateFunc)
}
