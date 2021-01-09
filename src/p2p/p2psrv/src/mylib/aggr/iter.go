package aggr

import (
	"sync"
)

func NewIteratorStu(Aggregate AggregateIntf) IterIntf {
	var Iter IteratorStu
	Iter.Init(Aggregate)
	return &Iter

}

type IteratorStu struct {
	Aggregate AggregateIntf
	Idx       int
	sync.RWMutex
}

func (r *IteratorStu) Init(Aggregate AggregateIntf) {
	r.Lock()
	defer r.Unlock()

	r.Aggregate = Aggregate
	r.Idx = 0
}

func (r *IteratorStu) HasNext() bool {
	r.Lock()
	defer r.Unlock()

	return r.Idx < r.Aggregate.Length()
}

func (r *IteratorStu) Next() (Token EntryIntf) {
	r.Lock()
	defer r.Unlock()

	if Token, ok := r.Aggregate.GetByIdx(r.Idx); ok {
		r.Idx++
		return Token
	}
	return nil

}
