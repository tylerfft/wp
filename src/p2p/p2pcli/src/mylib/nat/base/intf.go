package base

import (
	"net"
)

type EntryIntf interface {
	Init([]byte, *net.UDPAddr) int
	Execute() int
}
