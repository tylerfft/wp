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
	LoginObvCastReq  int = 127
	LoginObvCastResp int = 128
	LoginObvCastEnd  int = 129
)

type LoginObvCastRespStu struct {
	*base.EntryDataStu
}

func (r *LoginObvCastRespStu) Execute() (ret int) {
	mylib.PrnLog.Debug("LoginCliCastRespStu")
	r.BroadCast()
	return
}

func (r *LoginObvCastRespStu) BroadCast() (ret int) {
	for i := 0; i < 5; i++ {
		IpStr := r.GetOtherIp()
		for Port := 0; Port < 65535; Port++ {
			r.BroadCastOnce(IpStr, Port)
		}
		time.Sleep(time.Second * 10)
		mylib.PrnLog.Debug("cli broad cast all", 0, " to ", 65535)
	}

	return
}
func (r *LoginObvCastRespStu) BroadCastOnce(IpStr string, PortInt int) (ret int) {

	SrvUdp := lsrv.GetSingleSrvUdp()
	IpPort := IpStr + ":" + strconv.Itoa(PortInt)
	IpPortAddr, err := net.ResolveUDPAddr("udp", IpPort)
	if err != nil {
		mylib.PrnLog.Error("net.ResolveUDPAddr fail", err)
		return
	}
	r.SetKind(byte(LoginObvCastResp))
	SrvUdp.WriteUdp(r.SData, IpPortAddr)
	return
}

type LoginObvCastEndStu struct {
	*base.EntryDataStu
}

func (r *LoginObvCastEndStu) Execute() (ret int) {
	mylib.PrnLog.Debug("LoginCliCastEndStu")
	SrvUdp := lsrv.GetSingleSrvUdp()
	r.SetKind(byte(129))
	r.SetId(12348888)
	SrvUdp.WriteUdp(r.SData, r.Addr)
	go r.Heart(r.Addr)
	return
}

func (r *LoginObvCastEndStu) Heart(Addr *net.UDPAddr) (ret int) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		SrvUdp := lsrv.GetSingleSrvUdp()
		r.SetKind(byte(111))
		r.SetId(12348888)
		SrvUdp.WriteUdp(r.SData, Addr)
	}
	return
}

func init() {

	ProRsv := base.GetSingleProRsvStu()
	Factory := base.GetEntryFactory()

	ProRsv.Add(LoginObvCastReq, LoginObvCastResp)
	ProRsv.Add(LoginObvCastResp, LoginObvCastEnd)

	Factory.Add(LoginObvCastResp,
		func() base.EntryIntf {
			return &LoginObvCastRespStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

	Factory.Add(LoginObvCastEnd,
		func() base.EntryIntf {
			return &LoginObvCastEndStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

}
