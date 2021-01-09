package lsrv

import (
	mylib "mylib"
	base "mylib/nat/base"
)

var EntryVedioReq int = 234
var EntryVedioEnd int = 235

type VedioReqStu struct {
	*base.EntryDataStu
	Chan chan int
}

func (r *VedioReqStu) Execute() (ret int) {

	return
}

type VedioEndStu struct {
	*base.EntryDataStu
	Chan chan int
}

func (r *VedioEndStu) Execute() (ret int) {
	mylib.PrnLog.Debug("vidio", r.GetVedioStr())
	return
}

func init() {

	ProRsv := base.GetSingleProRsvStu()
	Factory := base.GetEntryFactory()

	ProRsv.Add(EntryVedioReq, EntryVedioEnd)

	Factory.Add(EntryVedioReq,
		func() base.EntryIntf {
			return &VedioReqStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

	Factory.Add(EntryVedioEnd,
		func() base.EntryIntf {
			return &VedioEndStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

}
