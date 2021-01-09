package nat

import (
	mylib "mylib"
	base "mylib/nat/base"
	mbase "p2p/nat/base"
)

var (
	KindCli02Obv01 string = base.NatStr[base.NATRestricted] + base.NatStr[base.NATFull]
	KindCli03Obv01 string = base.NatStr[base.NATPortRestricted] + base.NatStr[base.NATFull]
	KindCli04Obv01 string = base.NatStr[base.NATSymetric] + base.NatStr[base.NATFull]
)

func init() {

	Factory := mbase.GetMethodFactory()
	Factory.Add(KindCli02Obv01,
		func() mbase.EntryIntf {
			return &Cli04Obv01Stu{
				EntryStu: mbase.NewEntryStu(),
			}
		})

	Factory.Add(KindCli03Obv01,
		func() mbase.EntryIntf {
			return &Cli04Obv01Stu{
				EntryStu: mbase.NewEntryStu(),
			}
		})

	Factory.Add(KindCli04Obv01,
		func() mbase.EntryIntf {
			return &Cli04Obv01Stu{
				EntryStu: mbase.NewEntryStu(),
			}
		})
}

type Cli04Obv01Stu struct {
	*mbase.EntryStu
	Data *base.EntryDataStu
}

func (r *Cli04Obv01Stu) Execute() (ret int) {
	mylib.PrnLog.Debug("Cli04Obv01Stu")
	r.Data = base.NewEntryDataStu(nil, nil)
	r.Data.SetKind(byte(217))
	r.Data.SetOtherAddr(r.RAddr)
	r.WriteAddr(r.Data.SData, r.LAddr)
	return
}
