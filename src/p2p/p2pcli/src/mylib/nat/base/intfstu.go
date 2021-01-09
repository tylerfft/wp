package base

import (
	"net"
)

func NewEntryStu() *EntryStu {
	var v EntryStu
	return &v
}

type EntryStu struct {
	LAddr *net.UDPAddr
	RAddr *net.UDPAddr
	conn  *net.UDPConn
}

func (r *EntryStu) Init(lAddr *net.UDPAddr, rAddr *net.UDPAddr, conn *net.UDPConn) {
	r.LAddr = lAddr
	r.RAddr = rAddr
	r.conn = conn
	return
}
func (r *EntryStu) WriteAddr(Data []byte, Addr *net.UDPAddr) {
	r.conn.WriteTo(Data, Addr)
	return
}
