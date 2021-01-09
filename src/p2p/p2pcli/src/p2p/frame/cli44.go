package frame

import (
	"net"
	"strconv"

	"time"

	mylib "mylib"
	base "mylib/nat/base"
	lsrv "p2p/lsrv"
)

var (
	Cli44Req  int = 154
	Cli44Resp int = 155
	Cli44End  int = 156
)

type Cli44RespStu struct {
	*base.EntryDataStu
}

func (r *Cli44RespStu) Execute() (ret int) {
	mylib.PrnLog.Debug("Cli44RespStu")

	Port := r.GetOtherPortInt()
	IpStr := r.GetOtherIp()
	Port = Port + 1
	IpPort := IpStr + ":" + strconv.Itoa(Port)
	SrvUdp := lsrv.GetSingleSrvUdp()
	for i := 0; i < 30; i++ {

		IpPortAddr, err := net.ResolveUDPAddr("udp", IpPort)
		if err != nil {
			mylib.PrnLog.Error("net.ResolveUDPAddr fail", err)
			return
		}
		r.SetKind(byte(154))
		SrvUdp.WriteUdp(r.SData, IpPortAddr)

		mylib.PrnLog.Debug("trying to", IpPort)
		time.Sleep(time.Second * 1)
	}
	return
}

type Cli44EndStu struct {
	*base.EntryDataStu
}

func (r *Cli44EndStu) Execute() (ret int) {
	mylib.PrnLog.Debug("Cli44EndStu")
	SrvUdp := lsrv.GetSingleSrvUdp()
	r.SetKind(byte(208))
	r.SetId(12349999)
	SrvUdp.WriteUdp(r.SData, r.Addr)
	go r.Heart(r.Addr)
	return
}

func (r *Cli44EndStu) Heart(Addr *net.UDPAddr) (ret int) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		SrvUdp := lsrv.GetSingleSrvUdp()
		r.SetKind(byte(111))
		SrvUdp.WriteUdp(r.SData, Addr)
	}
	return
}

func init() {

	ProRsv := base.GetSingleProRsvStu()
	Factory := base.GetEntryFactory()

	ProRsv.Add(Cli44Req, Cli44Resp)
	ProRsv.Add(Cli44Resp, Cli44End)

	Factory.Add(Cli44Resp,
		func() base.EntryIntf {
			return &Cli44RespStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

	Factory.Add(Cli44End,
		func() base.EntryIntf {
			return &Cli44EndStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})
}
