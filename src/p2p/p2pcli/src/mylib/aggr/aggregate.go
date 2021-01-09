package aggr

import (
	"sync"
)

var SingleAggregate AggregateStu
var SingleAggregateOnce sync.Once

func GetSingleAggregate() AggregateIntf {
	SingleAggregateOnce.Do(SingleAggregate.Init)
	return &SingleAggregate
}
func NewAggregateStu() AggregateIntf {
	var Aggregate AggregateStu
	Aggregate.Init()
	return &Aggregate
}

type AggregateStu struct {
	Tokens []EntryIntf
	Valids []bool
	IdxMap map[string]int
	Len    int
	sync.RWMutex
}

func (r *AggregateStu) Init() {
	r.IdxMap = make(map[string]int)
}

func (r *AggregateStu) Add(Id string, Token EntryIntf) (Idx, ret int) {
	r.Lock()
	defer r.Unlock()

	idx, ok := r.IdxMap[Id]

	if ok {
		r.Tokens[idx] = Token
		r.Valids[idx] = true
		Idx = idx
		return
	}

	if r.Len < len(r.Tokens) {
		r.Tokens[r.Len] = Token
		r.Valids[r.Len] = true
		r.IdxMap[Id] = r.Len
		Idx = r.Len
	} else {
		r.Tokens = append(r.Tokens, Token)
		r.Valids = append(r.Valids, true)
		r.IdxMap[Id] = r.Len
		Idx = r.Len
	}
	r.Len++
	return
}
func (r *AggregateStu) Delete(IdDel string) {
	r.Lock()
	defer r.Unlock()

	if idxDel, iDelok := r.IdxMap[IdDel]; iDelok {
		IdLst := r.Tokens[r.Length()-1].GetId()
		r.IdxMap[IdLst] = idxDel
		r.Tokens[idxDel] = r.Tokens[r.Length()-1]

		r.Len--
	}
	delete(r.IdxMap, IdDel)
}

func (r *AggregateStu) Length() int {
	return r.Len
}
func (r *AggregateStu) GetByIdx(Idx int) (Token EntryIntf, ret bool) {
	r.RLock()
	defer r.RUnlock()

	if Idx >= r.Length() {
		return
	}
	if r.Valids[Idx] == false {
		return
	}
	Token = r.Tokens[Idx]
	ret = true
	return
}

func (r *AggregateStu) Find(Id string) (Token EntryIntf, exist bool) {
	r.RLock()
	defer r.RUnlock()
	if idx, ok := r.IdxMap[Id]; ok {
		Token = r.Tokens[idx]
		exist = true
	}
	return
}

func (r *AggregateStu) SetToken(Tokens []EntryIntf) {
	r.RLock()
	defer r.RUnlock()
	r.Tokens = Tokens
	r.Valids = make([]bool, len(r.Tokens))
	for k, v := range r.Tokens {
		r.Valids[k] = true
		r.IdxMap[v.GetId()] = k
	}
}

func (r *AggregateStu) Iterator() (Iter IterIntf) {
	r.RLock()
	defer r.RUnlock()
	return NewIteratorStu(r)
}
