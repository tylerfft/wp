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
	Obv44Req  int = 144
	Obv44Resp int = 145
	Obv44End  int = 146
)

type Obv44RespStu struct {
	*base.EntryDataStu
}

func (r *Obv44RespStu) Execute() (ret int) {
	mylib.PrnLog.Debug("Obv44RespStu")
	PortStart := r.GetOtherPortInt()
	IpStr := r.GetOtherIp()
	Port := PortStart + 2
	IpPort := IpStr + ":" + strconv.Itoa(Port)
	SrvUdp := lsrv.GetSingleSrvUdp()
	for i := 0; i < 5; i++ {
		IpPortAddr, err := net.ResolveUDPAddr("udp", IpPort)
		if err != nil {
			mylib.PrnLog.Error("net.ResolveUDPAddr fail", err)
			return
		}
		r.SetKind(byte(144))
		SrvUdp.WriteUdp(r.SData, IpPortAddr)

		mylib.PrnLog.Debug("trying to", IpPort)
		time.Sleep(time.Second * 5)
	}
	return
}

type Obv44EndStu struct {
	*base.EntryDataStu
}

func (r *Obv44EndStu) Execute() (ret int) {
	mylib.PrnLog.Debug("Obv44EndStu")
	SrvUdp := lsrv.GetSingleSrvUdp()
	r.SetKind(byte(208))
	r.SetId(12349999)
	SrvUdp.WriteUdp(r.SData, r.Addr)
	go r.Heart(r.Addr)
	return
}

func (r *Obv44EndStu) Heart(Addr *net.UDPAddr) (ret int) {
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

	ProRsv.Add(Obv44Req, Obv44Resp)
	ProRsv.Add(Obv44Resp, Obv44End)

	Factory.Add(Obv44Resp,
		func() base.EntryIntf {
			return &Obv44RespStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

	Factory.Add(Obv44End,
		func() base.EntryIntf {
			return &Obv44EndStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})
}
