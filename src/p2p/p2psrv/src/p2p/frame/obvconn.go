package frame

import (
	mylib "mylib"
	base "mylib/nat/base"
	clis "p2p/cli"
	nat "p2p/nat"
	natbase "p2p/nat/base"
	srv "p2p/srv"
)

var EntryObvConnReq int = 57
var EntryObvConnResp int = 58

type ObvConnRespStu struct {
	*base.EntryDataStu
	LData  *base.EntryDataStu
	Method natbase.EntryIntf
}

func (r *ObvConnRespStu) DealCli() (ret int) {
	CliIdStr := r.GetCliIdStr()
	mylib.PrnLog.Debug("cli id ", CliIdStr)
	Clis := clis.GetSingleClients()
	Cli, e := Clis.Find(CliIdStr)
	if !e {
		ret = -1
		return
	}
	r.LData = Cli
	return
}

func (r *ObvConnRespStu) Distribute() (ret int) {
	SrvUdp := srv.GetSingleSrvUdp()
	Factory := natbase.GetMethodFactory()
	if r.LData.GetAddr().IP.String() == r.GetAddr().IP.String() {
		r.Method = Factory.Create(nat.MethodSameNat)
	} else {
		r.Method = Factory.Create(r.LData.GetNatTypeStr() + r.GetNatTypeStr())
	}
	if r.Method == nil {
		mylib.PrnLog.Error(" no method ---- ")
		ret = -1
		return
	}
	r.PrintMethod()
	r.Method.Init(r.LData, r.EntryDataStu, SrvUdp.GetConn())
	ret = r.Method.Execute()
	if ret != 0 {
		mylib.PrnLog.Error(" r.	Method.Execute()")
	}
	return
}

func (r *ObvConnRespStu) PrintMethod() (ret int) {
	mylib.PrnLog.Debug(" method ---- ", r.LData.GetNatTypeStr())
	mylib.PrnLog.Debug(" method ---- ", r.LData.GetAddr())
	mylib.PrnLog.Debug(" method ---- ", r.GetNatTypeStr())
	mylib.PrnLog.Debug(" method ---- ", r.GetAddr())
	return
}

func (r *ObvConnRespStu) Execute() (ret int) {
	mylib.PrnLog.Debug("Execute")
	ret = r.DealCli()
	if ret != 0 {
		mylib.PrnLog.Error("r.DealCli()")
	}

	ret = r.Distribute()
	if ret != 0 {
		mylib.PrnLog.Error("r.Distribute()")
	}
	return
}

func init() {

	ProRsv := base.GetSingleProRsvStu()
	Factory := base.GetEntryFactory()

	ProRsv.Add(EntryObvConnReq, EntryObvConnResp)

	Factory.Add(EntryObvConnResp,
		func() base.EntryIntf {
			return &ObvConnRespStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

}
