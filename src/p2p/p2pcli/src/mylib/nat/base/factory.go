package base

import (
	"sync"
)

func NewEntry(Kind int) (Entry EntryIntf) {
	CreateFunc := GetEntryFactory()
	Entry = CreateFunc.Create(Kind)
	return
}

var SingleEntryFactory EntryFactoryStu
var SingleEntryFactoryOnce sync.Once

func GetEntryFactory() *EntryFactoryStu {
	SingleEntryFactoryOnce.Do(SingleEntryFactory.Init)
	return &SingleEntryFactory
}

type EntryCreateFunc func() (Entry EntryIntf)

type EntryFactoryStu struct {
	Factorys       map[int]EntryCreateFunc
	UnknownFactory EntryCreateFunc
}

func (r *EntryFactoryStu) Create(Kind int) (Entry EntryIntf) {
	if CreateFunc, ok := r.Factorys[Kind]; ok {
		Entry = CreateFunc()
	} else {
		Entry = nil
	}
	return
}

func (r *EntryFactoryStu) Add(Kind int, CreateFunc EntryCreateFunc) {
	r.Factorys[Kind] = CreateFunc
}

func (r *EntryFactoryStu) Init() {
	r.Factorys = make(map[int]EntryCreateFunc)
}
