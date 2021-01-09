package base

import (
	"net"

	mylib "mylib"
)

func Facade(Data []byte, Addr *net.UDPAddr) (ret int) {

	if len(Data) < 5 {
		mylib.PrnLog.Debug("")
		return
	}
	ProRsv := GetSingleProRsvStu()
	Factory := GetEntryFactory()
	Pro := int(Data[KindStart])
	Rsv := ProRsv.GetRsv(Pro)
	mylib.PrnLog.Debug("pro ---> ", Pro)
	mylib.PrnLog.Debug("rsv ---> ", Rsv)
	if Rsv == 0 {
		return
	}

	Entry := Factory.Create(Rsv)

	if Entry == nil {
		mylib.PrnLog.Error("Entry == nil", Data)
		mylib.PrnLog.Error("Entry == nil", string(Data))
		return
	}
	ret = Entry.Init(Data, Addr)
	if ret != 0 {
		mylib.PrnLog.Error("Entry.Initialize")
		return
	}
	ret = Entry.Execute()
	if ret != 0 {
		return
	}
	return
}
