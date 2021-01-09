package nat

import (
	mylib "mylib"
	base "mylib/nat/base"
	mbase "p2p/nat/base"
)

var (
	KindCli03Obv04 string = base.NatStr[base.NATPortRestricted] + base.NatStr[base.NATSymetric]
)

func init() {

	Factory := mbase.GetMethodFactory()
	Factory.Add(KindCli03Obv04,
		func() mbase.EntryIntf {
			return &Cli03Obv04Stu{
				EntryStu: mbase.NewEntryStu(),
			}
		})
}

type Cli03Obv04Stu struct {
	*mbase.EntryStu
	Data01 *base.EntryDataStu
	Data02 *base.EntryDataStu
}

func (r *Cli03Obv04Stu) Execute() (ret int) {
	mylib.PrnLog.Debug("Cli03Obv04Stu ------------------------------ ")

	r.Data01 = base.NewEntryDataStu(nil, nil)
	r.Data01.SetKind(byte(137))
	r.Data01.SetOtherAddr(r.RAddr)
	r.WriteAddr(r.Data01.SData, r.LAddr)

	r.Data02 = base.NewEntryDataStu(nil, nil)
	r.Data02.SetKind(byte(127))
	r.Data02.SetOtherAddr(r.LAddr)
	r.WriteAddr(r.Data02.SData, r.RAddr)
	return
}
