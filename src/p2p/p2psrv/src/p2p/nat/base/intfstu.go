package base

import (
	"net"

	base "mylib/nat/base"
)

func NewEntryStu() *EntryStu {
	var v EntryStu
	return &v
}

type EntryStu struct {
	LData *base.EntryDataStu
	RData *base.EntryDataStu
	LAddr *net.UDPAddr
	RAddr *net.UDPAddr
	conn  *net.UDPConn
}

func (r *EntryStu) Init(lData *base.EntryDataStu, rData *base.EntryDataStu, conn *net.UDPConn) {
	r.LData = lData
	r.RData = rData
	r.LAddr = r.LData.GetAddr()
	r.RAddr = r.RData.GetAddr()
	r.conn = conn
	return
}
func (r *EntryStu) WriteAddr(Data []byte, Addr *net.UDPAddr) {
	r.conn.WriteTo(Data, Addr)
	return
}
