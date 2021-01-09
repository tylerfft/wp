package base

import (
	"net"

	base "mylib/nat/base"
)

type EntryIntf interface {
	Init(*base.EntryDataStu, *base.EntryDataStu, *net.UDPConn)
	Execute() int
}
