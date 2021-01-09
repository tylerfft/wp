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
	LoginCliCastReq  int = 227
	LoginCliCastResp int = 228
	LoginCliCastEnd  int = 229
)

type LoginCliCastRespStu struct {
	*base.EntryDataStu
}

func (r *LoginCliCastRespStu) Execute() (ret int) {
	mylib.PrnLog.Debug("LoginCliCastRespStu")
	r.BroadCast()
	return
}

func (r *LoginCliCastRespStu) BroadCast() (ret int) {
	IpStr := r.GetOtherIp()

	for i := 0; i < 54; i++ {
		for Port := 0; Port < 65535; Port++ {
			r.BroadCastOnce(IpStr, Port)
		}
		time.Sleep(time.Second * 10)
	}
	mylib.PrnLog.Debug("cli broad cast all ", 0, " to ", 65535)
	return
}
func (r *LoginCliCastRespStu) BroadCastOnce(IpStr string, PortInt int) (ret int) {

	SrvUdp := lsrv.GetSingleSrvUdp()
	IpPort := IpStr + ":" + strconv.Itoa(PortInt)
	IpPortAddr, err := net.ResolveUDPAddr("udp", IpPort)
	if err != nil {
		mylib.PrnLog.Error("net.ResolveUDPAddr fail", err)
		return
	}
	r.SetKind(byte(LoginCliCastResp))

	SrvUdp.WriteUdp(r.SData, IpPortAddr)
	return
}

type LoginCliCastEndStu struct {
	*base.EntryDataStu
}

func (r *LoginCliCastEndStu) Execute() (ret int) {
	for i := 0; i < 3; i++ {
		mylib.PrnLog.Debug("LoginCliCastEndStu")
		SrvUdp := lsrv.GetSingleSrvUdp()
		r.SetKind(byte(208))
		SrvUdp.WriteUdp(r.SData, r.Addr)
		time.Sleep(time.Second * 5)
	}
	go r.Heart(r.Addr)
	return
}

func (r *LoginCliCastEndStu) Heart(Addr *net.UDPAddr) (ret int) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		SrvUdp := lsrv.GetSingleSrvUdp()
		r.SetKind(byte(111))
		r.SetId(12349999)
		SrvUdp.WriteUdp(r.SData, Addr)
	}
	return
}

func init() {

	ProRsv := base.GetSingleProRsvStu()
	Factory := base.GetEntryFactory()

	ProRsv.Add(LoginCliCastReq, LoginCliCastResp)
	ProRsv.Add(LoginCliCastResp, LoginCliCastEnd)

	Factory.Add(LoginCliCastResp,
		func() base.EntryIntf {
			return &LoginCliCastRespStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

	Factory.Add(LoginCliCastEnd,
		func() base.EntryIntf {
			return &LoginCliCastEndStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})
}
