package obv

import (
	"fmt"
	"sync"
	"time"

	aggr "mylib/aggr"
	base "mylib/nat/base"
	lsrv "p2p/lsrv"
)

var SingleObservers ObserversStu
var SingleObserversOnce sync.Once

func GetSingleObservers() *ObserversStu {
	SingleObserversOnce.Do(SingleObservers.Init)
	return &SingleObservers
}

type ObserversStu struct {
	aggr aggr.AggregateIntf
}

func (r *ObserversStu) Add(cli *base.EntryDataStu) {
	r.aggr.Add(cli.GetId(), cli)
}
func (r *ObserversStu) Find(Id string) (cli *base.EntryDataStu, exist bool) {
	if v, e := r.aggr.Find(Id); e {
		if vv, ok := v.(*base.EntryDataStu); ok {
			cli = vv
			exist = true
		}
	}
	return
}
func (r *ObserversStu) Init() {
	r.aggr = aggr.NewAggregateStu()
}
func (r *ObserversStu) ShowObservers() {
	ticker := time.NewTicker(time.Second * 5)
	Lsrv := lsrv.GetSingleSrvUdp()
	for range ticker.C {
		Iter := r.aggr.Iterator()
		fmt.Println("obvs :", "--------")
		for Iter.HasNext() {
			v := Iter.Next()
			if vv, ok := v.(*base.EntryDataStu); ok {
				fmt.Println(vv.GetId(), " @ ", vv.Addr)
				Data := base.NewEntryDataStu(nil, nil)
				Data.SetKind(byte(234))
				Data.SetVedioStr("i am camera")
				go Lsrv.WriteUdp(Data.SData, vv.Addr)
			}
		}
	}
}

func init() {
	Obvs := GetSingleObservers()
	go Obvs.ShowObservers()
}
