package cli

import (
	"sync"

	aggr "mylib/aggr"
	base "mylib/nat/base"
)

var SingleClients ClientsStu
var SingleClientsOnce sync.Once

func GetSingleClients() *ClientsStu {
	SingleClientsOnce.Do(SingleClients.Init)
	return &SingleClients
}

type ClientsStu struct {
	Aggr aggr.AggregateIntf
}

func (r *ClientsStu) Add(cli *base.EntryDataStu) {
	r.Aggr.Add(cli.GetId(), cli)
}
func (r *ClientsStu) Find(Id string) (out *base.EntryDataStu, exist bool) {
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

func (r *ClientsStu) Init() {
	r.Aggr = aggr.NewAggregateStu()
}
