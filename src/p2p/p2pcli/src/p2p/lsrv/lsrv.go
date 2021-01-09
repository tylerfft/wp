package lsrv

import (
	"net"
	"os"
	"sync"
	"time"

	mylib "mylib"
	base "mylib/nat/base"
	setting "p2p/setting"

	stun "github.com/ccding/go-stun/stun"
)

var SingleSrvUdp SrvUdpStu
var SingleSrvUdpOnce sync.Once

func GetSingleSrvUdp() *SrvUdpStu {
	SingleSrvUdpOnce.Do(SingleSrvUdp.Init)
	return &SingleSrvUdp
}

type SrvUdpStu struct {
	Conn    *net.UDPConn
	Addr    *net.UDPAddr
	Cli     *stun.Client
	NatType int
}

func (r *SrvUdpStu) Init() {
	var err error

	r.Addr, err = net.ResolveUDPAddr("udp", ":0")
	if err != nil {
		mylib.PrnLog.Error("Can't resolve address: ", err)
		os.Exit(1)
	}
	r.Conn, err = net.ListenUDP("udp", r.Addr)
	if err != nil {
		mylib.PrnLog.Error("net.ListenUDP Addr ---", r.Addr)
		mylib.PrnLog.Error("net.ListenUDP err ---", err)
		return
	}
	t := time.Now()
	r.Conn.SetDeadline(t.Add(time.Duration(30 * time.Second)))
	r.Cli = stun.NewClientWithConnection(r.Conn)
	return
}
func (r *SrvUdpStu) Login() (ret int) {

	a, b, err := r.Cli.Discover()
	if err != nil {
		mylib.PrnLog.Error("err", err)
		return
	}
	mylib.PrnLog.Debug("a", a)
	mylib.PrnLog.Debug("b", b)

	LoginReq := GetSingleLoginReq()

	Setting := setting.GetSetting()
	if Setting.Model == "obv" {
		LoginReq.SetKind(byte(67))
	} else {
		LoginReq.SetKind(byte(77))
	}

	LoginReq.SetId(12348888)
	LoginReq.SetNatType(byte(a))
	r.NatType = int(a)
	LoginReq.SetRole(byte(1))
	go LoginReq.Execute()

	return
}
func (r *SrvUdpStu) Read() {
	for {
		data := make([]byte, 1024)
		n, addr, err := r.Conn.ReadFromUDP(data)
		if err != nil {
			mylib.PrnLog.Error("err", err)
			time.Sleep(time.Second * 1)
			continue
		} else {
			mylib.PrnLog.Debug("addr", addr)
			go base.Facade(data[:n], addr)
		}
	}
}
func (r *SrvUdpStu) ReadTicker() {
	ticker := time.NewTicker(time.Millisecond * 5)
	for range ticker.C {
		t := time.Now()
		r.Conn.SetDeadline(t.Add(time.Duration(5 * time.Second)))
	}
}
func (r *SrvUdpStu) Heart2srv() {
	raddrstr := "180.76.119.248:55500"
	raddr, err := net.ResolveUDPAddr("udp", raddrstr)
	if err != nil {
		mylib.PrnLog.Error("net.ResolveUDPAddr", err)
		os.Exit(1)
	}
	ticker := time.NewTicker(time.Second * 20)
	for range ticker.C {
		r.Conn.WriteTo([]byte("Hello"), raddr)
	}
}
func (r *SrvUdpStu) WriteUdp(Data []byte, Addr net.Addr) {
	r.Conn.WriteTo(Data, Addr)
}
func init() {
	SrvUdp := GetSingleSrvUdp()

	ret := SrvUdp.Login()
	if ret != 0 {
		mylib.PrnLog.Error("Login")
	}
	go SrvUdp.Read()
	go SrvUdp.ReadTicker()
	go SrvUdp.Heart2srv()

}
