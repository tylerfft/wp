package frame

import (
	"net"
	"time"

	mylib "mylib"
	base "mylib/nat/base"
	lsrv "p2p/lsrv"
)

var (
	EntryOtherLoginCmdReq  int = 217
	EntryOtherLoginCmdResp int = 218
	EntryOtherLoginCmdEnd  int = 219
)

type EntryOtherLoginCmdRespStu struct {
	*base.EntryDataStu
}

func (r *EntryOtherLoginCmdRespStu) Execute() (ret int) {
	mylib.PrnLog.Debug(" EntryOtherLoginCmdRespStu")
	SrvUdp := lsrv.GetSingleSrvUdp()
	OtherAddr := r.GetOtherAddr()
	mylib.PrnLog.Debug(" OtherAddr", OtherAddr)
	r.SetKind(byte(EntryOtherLoginCmdResp))
	SrvUdp.WriteUdp(r.SData, OtherAddr)
	go r.Heart(OtherAddr)
	return
}
func (r *EntryOtherLoginCmdRespStu) Heart(Addr *net.UDPAddr) (ret int) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		SrvUdp := lsrv.GetSingleSrvUdp()
		r.SetKind(byte(111))
		r.SetId(12349999)
		SrvUdp.WriteUdp(r.SData, Addr)
	}
	return
}

type EntryOtherLoginCmdEndStu struct {
	*base.EntryDataStu
}

func (r *EntryOtherLoginCmdEndStu) Execute() (ret int) {
	mylib.PrnLog.Debug(" EntryOtherLoginCmdEndStu")
	SrvUdp := lsrv.GetSingleSrvUdp()
	r.SetKind(byte(208))
	r.SetId(12349999)
	SrvUdp.WriteUdp(r.SData, r.Addr)
	go r.Heart(r.Addr)
	return
}

func (r *EntryOtherLoginCmdEndStu) Heart(Addr *net.UDPAddr) (ret int) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		SrvUdp := lsrv.GetSingleSrvUdp()
		r.SetKind(byte(111))
		r.SetId(12349999)
		SrvUdp.WriteUdp(r.SData, Addr)
	}
	return
}

func init() {

	ProRsv := base.GetSingleProRsvStu()
	Factory := base.GetEntryFactory()

	ProRsv.Add(EntryOtherLoginCmdReq, EntryOtherLoginCmdResp)
	ProRsv.Add(EntryOtherLoginCmdResp, EntryOtherLoginCmdEnd)

	Factory.Add(EntryOtherLoginCmdResp,
		func() base.EntryIntf {
			return &EntryOtherLoginCmdRespStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

	Factory.Add(EntryOtherLoginCmdEnd,
		func() base.EntryIntf {
			return &EntryOtherLoginCmdEndStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})
}
