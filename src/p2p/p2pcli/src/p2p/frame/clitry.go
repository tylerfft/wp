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
	LoginCliTrySrv2One int = 137
	LoginCliTryOne2Oth int = 138
	LoginCliTryOth2One int = 139
	LoginCliTryOneEnd  int = 140
)

type LoginCliTryOne2OthStu struct {
	*base.EntryDataStu
}

func (r *LoginCliTryOne2OthStu) Execute() (ret int) {
	mylib.PrnLog.Debug("LoginCliTryOne2OthStu")
	PortStart := r.GetOtherPortInt()
	IpStr := r.GetOtherIp()
	SrvUdp := lsrv.GetSingleSrvUdp()
	for i := 0; i < 10; i++ {
		for Port := 0; Port < 65535; Port++ {
			IpPort := IpStr + ":" + strconv.Itoa(Port)
			IpPortAddr, err := net.ResolveUDPAddr("udp", IpPort)
			if err != nil {
				mylib.PrnLog.Error("net.ResolveUDPAddr fail", err)
				return
			}
			r.SetKind(byte(138))
			SrvUdp.WriteUdp(r.SData, IpPortAddr)
		}
		mylib.PrnLog.Debug("trying to", IpStr)
		mylib.PrnLog.Debug(" from ", PortStart, " to ", 65535)
		time.Sleep(time.Second * 5)
	}
	return
}

type LoginCliTryOth2OneStu struct {
	*base.EntryDataStu
}

func (r *LoginCliTryOth2OneStu) Execute() (ret int) {
	mylib.PrnLog.Debug("LoginTryOth2OneStu")
	r.SetKind(byte(208))
	r.SetId(12348888)
	SrvUdp := lsrv.GetSingleSrvUdp()
	SrvUdp.WriteUdp(r.SData, r.Addr)
	go r.Heart()
	return
}
func (r *LoginCliTryOth2OneStu) Heart() (ret int) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		SrvUdp := lsrv.GetSingleSrvUdp()
		r.SetKind(byte(111))
		SrvUdp.WriteUdp(r.SData, r.Addr)
	}
	return
}

type LoginCliTryOneEndStu struct {
	*base.EntryDataStu
}

func (r *LoginCliTryOneEndStu) Execute() (ret int) {
	mylib.PrnLog.Debug("LoginTryOneEndStu")

	return
}

type LoginCliCastDoneStu struct {
	*base.EntryDataStu
}

func (r *LoginCliCastDoneStu) Execute() (ret int) {
	mylib.PrnLog.Debug("LoginCliCastEndStu")
	SrvUdp := lsrv.GetSingleSrvUdp()
	r.SetKind(byte(208))
	r.SetId(12348888)
	SrvUdp.WriteUdp(r.SData, r.Addr)
	go r.Heart(r.Addr)
	return
}
func (r *LoginCliCastDoneStu) Heart(Addr *net.UDPAddr) (ret int) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		SrvUdp := lsrv.GetSingleSrvUdp()
		r.SetKind(byte(111))
		r.SetId(12348888)
		SrvUdp.WriteUdp(r.SData, Addr)
	}
	return
}

var SingleCliTryChan chan int

func init() {

	SingleCliTryChan = make(chan int)

	ProRsv := base.GetSingleProRsvStu()
	Factory := base.GetEntryFactory()

	ProRsv.Add(LoginCliTrySrv2One, LoginCliTryOne2Oth)
	ProRsv.Add(LoginCliTryOne2Oth, LoginCliTryOth2One)
	ProRsv.Add(LoginCliTryOth2One, LoginCliTryOneEnd)

	Factory.Add(LoginCliTryOne2Oth,
		func() base.EntryIntf {
			return &LoginCliTryOne2OthStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})

	Factory.Add(LoginCliTryOth2One,
		func() base.EntryIntf {
			return &LoginCliTryOth2OneStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})
	Factory.Add(LoginCliTryOneEnd,
		func() base.EntryIntf {
			return &LoginCliTryOneEndStu{
				EntryDataStu: base.NewEntryDataStu(nil, nil),
			}
		})
}
