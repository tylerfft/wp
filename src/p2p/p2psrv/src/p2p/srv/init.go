package srv

import (
	"fmt"
	"time"

	base "mylib/nat/base"
	clis "p2p/cli"
	obvs "p2p/obv"
)

func init() {
	SrvInit()
	go ScanObvs()
	go ScanClis()
}
func SrvInit() {
	SrvUdp := GetSingleSrvUdp()
	go SrvUdp.KeepAlive()
	fmt.Println("SrvUdp.Dis()")
	//	SrvUdp.Dis()
	go SrvUdp.Read()
}

func ScanObvs() {
	ticker := time.NewTicker(time.Second * 5)
	Obvs := obvs.GetSingleObservers()
	SrvUdp := GetSingleSrvUdp()
	for range ticker.C {
		Iter := Obvs.Aggr.Iterator()
		fmt.Println("obvs :", "--------")
		for Iter.HasNext() {
			v := Iter.Next()
			if vv, ok := v.(*base.EntryDataStu); ok {
				fmt.Println(vv.GetId(), " @ ", vv.Addr)
				SrvUdp.WriteAddr([]byte("i am server ---heart"), vv.Addr)
			}
		}
		fmt.Println("------------------------")
	}
}

func ScanClis() {
	ticker := time.NewTicker(time.Second * 5)
	Clis := clis.GetSingleClients()
	SrvUdp := GetSingleSrvUdp()
	for range ticker.C {
		Iter := Clis.Aggr.Iterator()
		fmt.Println("clis :", "--------")
		for Iter.HasNext() {
			v := Iter.Next()
			if vv, ok := v.(*base.EntryDataStu); ok {
				fmt.Println(vv.GetId(), " @ ", vv.Addr)
				SrvUdp.WriteAddr([]byte("i am server ---heart"), vv.Addr)
			}
		}
		fmt.Println("------------------------")
	}
}
