package nat

import (
	mylib "mylib"
	base "mylib/nat/base"
	mbase "p2p/nat/base"
)

var (
	KindCli04Obv03 string = base.NatStr[base.NATSymetric] + base.NatStr[base.NATPortRestricted]
)

func init() {

	Factory := mbase.GetMethodFactory()
	Factory.Add(KindCli04Obv03,
		func() mbase.EntryIntf {
			return &Cli04Obv03Stu{
				EntryStu: mbase.NewEntryStu(),
			}
		})
}

type Cli04Obv03Stu struct {
	*mbase.EntryStu
	Data01 *base.EntryDataStu
	Data02 *base.EntryDataStu
}

func (r *Cli04Obv03Stu) Execute() (ret int) {
	mylib.PrnLog.Debug("Cli04Obv03Stu ------------------------------ ")

	r.Data01 = base.NewEntryDataStu(nil, nil)
	r.Data01.SetKind(byte(227))
	r.Data01.SetOtherAddr(r.RAddr)
	r.WriteAddr(r.Data01.SData, r.LAddr)

	r.Data02 = base.NewEntryDataStu(nil, nil)
	r.Data02.SetKind(byte(237))
	r.Data02.SetOtherAddr(r.LAddr)
	r.WriteAddr(r.Data02.SData, r.RAddr)
	return
}
