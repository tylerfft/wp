package nat

import (
	mylib "mylib"
	base "mylib/nat/base"
	mbase "p2p/nat/base"
)

var (
	KindCli04Obv04 string = base.NatStr[base.NATSymetric] + base.NatStr[base.NATSymetric]
)

func init() {

	Factory := mbase.GetMethodFactory()
	Factory.Add(KindCli04Obv04,
		func() mbase.EntryIntf {
			return &Cli04Obv04Stu{
				EntryStu: mbase.NewEntryStu(),
			}
		})
}

type Cli04Obv04Stu struct {
	*mbase.EntryStu
	Data01 *base.EntryDataStu
	Data02 *base.EntryDataStu
}

func (r *Cli04Obv04Stu) Execute() (ret int) {
	mylib.PrnLog.Debug("Cli04Obv04Stu ------------------------------ ")

	r.Data01 = base.NewEntryDataStu(nil, nil)
	r.Data01.SetKind(byte(127))
	r.Data01.SetOtherAddr(r.RAddr)
	r.WriteAddr(r.Data01.SData, r.LAddr)
	r.Data01 = base.NewEntryDataStu(nil, nil)
	r.Data01.SetKind(byte(137))
	r.Data01.SetOtherAddr(r.RAddr)
	r.WriteAddr(r.Data01.SData, r.LAddr)

	r.Data02 = base.NewEntryDataStu(nil, nil)
	r.Data02.SetKind(byte(127))
	r.Data02.SetOtherAddr(r.LAddr)
	r.WriteAddr(r.Data02.SData, r.RAddr)
	r.Data02 = base.NewEntryDataStu(nil, nil)

	return
}
