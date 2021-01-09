package obv

import (
	"sync"

	aggr "mylib/aggr"
	base "mylib/nat/base"
)

var SingleObservers ObserversStu
var SingleObserversOnce sync.Once

func GetSingleObservers() *ObserversStu {
	SingleObserversOnce.Do(SingleObservers.Init)
	return &SingleObservers
}

type ObserversStu struct {
	Aggr aggr.AggregateIntf
}

func (r *ObserversStu) Add(v *base.EntryDataStu) {
	r.Aggr.Add(v.GetId(), v)
}
func (r *ObserversStu) Find(Id string) (out *base.EntryDataStu, exist bool) {
	v, e := r.Aggr.Find(Id)
	if e {
		if vv, ok := v.(*base.EntryDataStu); ok {
			out = vv
			exist = true
			return
		}
	}
	return
}
func (r *ObserversStu) Init() {
	r.Aggr = aggr.NewAggregateStu()
}
