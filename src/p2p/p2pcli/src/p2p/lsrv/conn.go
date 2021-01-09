package lsrv

import (
	"sync"

	mylib "mylib"
	base "mylib/nat/base"
	setting "p2p/setting"
)

var EntryObvConnReq int = 57
var EntryObvConnResp int = 58
var EntryObvConnEnd int = 59

var SingleObvConnReq ObvConnReqStu
var SingleObvConnReqOnce sync.Once

func GetSingleObvConnReq() *ObvConnReqStu {
	SingleObvConnReqOnce.Do(
		func() {
			SingleObvConnReq.EntryDataStu = base.NewEntryDataStu(nil, nil)
			SingleObvConnReq.SetKind(byte(57))
			SingleObvConnReq.SetCliId(12348888)
			SrvUdp := GetSingleSrvUdp()
			SingleObvConnReq.SetNatType(byte(SrvUdp.NatType))
		})

	return &SingleObvConnReq
}

type ObvConnReqStu struct {
	*base.EntryDataStu
}

func (r *ObvConnReqStu) Execute() (ret int) {
	mylib.PrnLog.Debug("LoginReqStu Execute")
	SrvUdp := GetSingleSrvUdp()
	Setting := setting.GetSetting()
	SrvUdp.WriteUdp(r.SData, Setting.GetSrvAddr())
	mylib.PrnLog.Debug("time")

	return
}

type ObvConnEndStu struct {
	*base.EntryDataStu
}

func (r *ObvConnEndStu) Execute() (ret int) {
	mylib.PrnLog.Debug("ObvConnEndStu")
	return
}

func init() {

	ProRsv := base.GetSingleProRsvStu()
	Factory := base.GetEntryFactory()

	ProRsv.Add(EntryCliLoginReq, EntryCliLoginResp)
	ProRsv.Add(EntryCliLoginResp, EntryCliLoginEnd)

	ProRsv.Add(EntryObvLoginReq, EntryObvLoginResp)
	ProRsv.Add(EntryObvLoginResp, EntryObvLoginEnd)

	Factory.Add(EntryObvConnReq,
		func() base.EntryIntf {
			return &ObvConnReqStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

	Factory.Add(EntryObvConnEnd,
		func() base.EntryIntf {
			return &ObvConnEndStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

}
