package frame

import (
	"net"
	"time"

	mylib "mylib"
	base "mylib/nat/base"
	lsrv "p2p/lsrv"
	obv "p2p/obv"
)

var (
	EntryOtherLoginReq  int = 207
	EntryOtherLoginResp int = 208
	EntryOtherLoginEnd  int = 209
	EntryOtherLoginDone int = 210
)

type EntryOtherLoginRespStu struct {
	*base.EntryDataStu
}

func (r *EntryOtherLoginRespStu) Execute() (ret int) {
	for i := 0; i < 3; i++ {
		mylib.PrnLog.Debug(" EntryOtherLoginRespStu")
		SrvUdp := lsrv.GetSingleSrvUdp()
		OtherAddr := r.GetOtherAddr()
		mylib.PrnLog.Debug(" OtherAddr", OtherAddr)
		r.SetKind(byte(EntryOtherLoginResp))
		r.SetId(12348888)
		SrvUdp.WriteUdp(r.SData, OtherAddr)
		time.Sleep(time.Second * 1)
	}
	return
}

type EntryOtherLoginEndStu struct {
	*base.EntryDataStu
}

func (r *EntryOtherLoginEndStu) Execute() (ret int) {
	mylib.PrnLog.Debug(" EntryOtherLoginEndStu")
	Obvs := obv.GetSingleObservers()
	Obv := base.NewEntryDataStu(r.RData, r.Addr)
	if _, e := Obvs.Find(Obv.GetId()); e {
		return
	}
	Obvs.Add(Obv)
	mylib.PrnLog.Debug("cli Add obvs", Obv.GetId())
	go r.Heart(r.Addr)
	return
}

func (r *EntryOtherLoginEndStu) Heart(Addr *net.UDPAddr) (ret int) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		SrvUdp := lsrv.GetSingleSrvUdp()
		r.SetKind(byte(111))

		SrvUdp.WriteUdp(r.SData, Addr)
	}
	return
}

func init() {

	ProRsv := base.GetSingleProRsvStu()
	Factory := base.GetEntryFactory()

	ProRsv.Add(EntryOtherLoginReq, EntryOtherLoginResp)
	ProRsv.Add(EntryOtherLoginResp, EntryOtherLoginEnd)

	Factory.Add(EntryOtherLoginResp,
		func() base.EntryIntf {
			return &EntryOtherLoginRespStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

	Factory.Add(EntryOtherLoginEnd,
		func() base.EntryIntf {
			return &EntryOtherLoginEndStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

}
