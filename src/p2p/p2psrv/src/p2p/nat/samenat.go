package nat

import (
	mylib "mylib"
	base "mylib/nat/base"
	mbase "p2p/nat/base"
)

var (
	MethodSameNat string = "MethodSameNat"
)

func init() {

	Factory := mbase.GetMethodFactory()
	Factory.Add(MethodSameNat,
		func() mbase.EntryIntf {
			return &SameNatStu{
				EntryStu: mbase.NewEntryStu(),
			}
		})
}

type SameNatStu struct {
	*mbase.EntryStu
	Data01 *base.EntryDataStu
	Data02 *base.EntryDataStu
}

func (r *SameNatStu) Execute() (ret int) {

	mylib.PrnLog.Debug("SameNatStu ------------------------------ ")
	if r.EntryStu.LData.GetNatTypeStr() == base.NatStr[base.NATSymetric] {
		mylib.PrnLog.Debug("Type4")
		r.Type123()
	} else {
		mylib.PrnLog.Debug("Type123")
		r.Type123()
	}

	return
}
func (r *SameNatStu) Type4() (ret int) {
	mylib.PrnLog.Debug("Type4")

	r.Data01 = base.NewEntryDataStu(nil, nil)
	r.Data01.SetKind(byte(144))
	r.Data01.SetOtherAddr(r.LAddr)
	r.WriteAddr(r.Data01.SData, r.RAddr)

	r.Data02 = base.NewEntryDataStu(nil, nil)
	r.Data02.SetKind(byte(154))
	r.Data02.SetOtherAddr(r.RAddr)
	r.WriteAddr(r.Data02.SData, r.LAddr)

	return
}

func (r *SameNatStu) Type123() (ret int) {
	mylib.PrnLog.Debug("Type123")

	r.Data01 = base.NewEntryDataStu(nil, nil)
	r.Data01.SetKind(byte(207))
	r.Data01.SetOtherAddr(r.LAddr)
	r.WriteAddr(r.Data01.SData, r.RAddr)

	r.Data02 = base.NewEntryDataStu(nil, nil)
	r.Data02.SetKind(byte(207))
	r.Data02.SetOtherAddr(r.RAddr)
	r.WriteAddr(r.Data02.SData, r.LAddr)

	return
}
