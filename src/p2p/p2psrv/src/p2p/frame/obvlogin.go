package frame

import (
	mylib "mylib"
	base "mylib/nat/base"
	obvs "p2p/obv"
	srv "p2p/srv"
)

var EntryObvLoginReq int = 67
var EntryObvLoginResp int = 68
var EntryObvLoginEnd int = 69

type ObvLoginRespStu struct {
	*base.EntryDataStu
	Chan chan int
}

func (r *ObvLoginRespStu) Execute() (ret int) {
	SrvUdp := srv.GetSingleSrvUdp()
	r.SetKind(byte(EntryObvLoginResp))
	mylib.PrnLog.Debug("r.String()", r.String())
	SrvUdp.WriteAddr(r.SData, r.Addr)

	Obvs := obvs.GetSingleObservers()
	Obv := base.NewEntryDataStu(r.RData, r.Addr)
	Obvs.Add(Obv)
	mylib.PrnLog.Debug("add obv", Obv.GetId())
	return
}

func init() {

	ProRsv := base.GetSingleProRsvStu()
	Factory := base.GetEntryFactory()

	ProRsv.Add(EntryObvLoginReq, EntryObvLoginResp)
	ProRsv.Add(EntryObvLoginResp, EntryObvLoginEnd)

	Factory.Add(EntryObvLoginResp,
		func() base.EntryIntf {
			return &ObvLoginRespStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

}
