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
	LoginObvTrySrv2One int = 237
	LoginObvTryOne2Oth int = 238
	LoginObvTryOth2One int = 239
	LoginObvTryOneEnd  int = 240
)

type LoginObvTryOne2OthStu struct {
	*base.EntryDataStu
}

func (r *LoginObvTryOne2OthStu) Execute() (ret int) {
	mylib.PrnLog.Debug("LoginObvTryOne2OthStu")
	//	PortStart := r.GetOtherPortInt()
	IpStr := r.GetOtherIp()
	SrvUdp := lsrv.GetSingleSrvUdp()
	for i := 0; i < 5; i++ {
		for Port := 0; Port < 65535; Port++ {
			IpPort := IpStr + ":" + strconv.Itoa(Port)
			IpPortAddr, err := net.ResolveUDPAddr("udp", IpPort)
			if err != nil {
				mylib.PrnLog.Error("net.ResolveUDPAddr fail", err)
				return
			}
			r.SetKind(byte(238))
			r.SetId(9999999)
			SrvUdp.WriteUdp(r.SData, IpPortAddr)
		}
		mylib.PrnLog.Debug("trying to", IpStr)
		mylib.PrnLog.Debug(" from ", 0, " to ", 65535)
		time.Sleep(time.Second * 5)
	}
	return
}

type LoginObvTryOth2OneStu struct {
	*base.EntryDataStu
}

func (r *LoginObvTryOth2OneStu) Execute() (ret int) {
	mylib.PrnLog.Debug("LoginTryOth2OneStu")
	SrvUdp := lsrv.GetSingleSrvUdp()
	r.SetKind(byte(239))
	r.SetId(99999999)
	SrvUdp.WriteUdp(r.SData, r.Addr)
	return
}

type LoginObvTryOneEndStu struct {
	*base.EntryDataStu
}

func (r *LoginObvTryOneEndStu) Execute() (ret int) {
	for i := 0; i < 10; i++ {
		mylib.PrnLog.Debug("LoginTryOneEndStu")
		SrvUdp := lsrv.GetSingleSrvUdp()
		r.SetKind(byte(208))
		r.SetId(12348888)
		SrvUdp.WriteUdp(r.SData, r.Addr)
		time.Sleep(time.Second * 5)
	}
	go r.Heart(r.Addr)
	return
}

func (r *LoginObvTryOneEndStu) Heart(Addr *net.UDPAddr) (ret int) {
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

	ProRsv.Add(LoginObvTrySrv2One, LoginObvTryOne2Oth)
	ProRsv.Add(LoginObvTryOne2Oth, LoginObvTryOth2One)
	ProRsv.Add(LoginObvTryOth2One, LoginObvTryOneEnd)

	Factory.Add(LoginObvTryOne2Oth,
		func() base.EntryIntf {
			return &LoginObvTryOne2OthStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

	Factory.Add(LoginObvTryOth2One,
		func() base.EntryIntf {
			return &LoginObvTryOth2OneStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})
	Factory.Add(LoginObvTryOneEnd,
		func() base.EntryIntf {
			return &LoginObvTryOneEndStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})
}
