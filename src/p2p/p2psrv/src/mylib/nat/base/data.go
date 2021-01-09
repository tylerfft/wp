package base

import (
	"fmt"
	"net"

	"strconv"

	mylib "mylib"
	tool "mylib/tool"
)

func NewEntryDataStu(Data []byte, Addr *net.UDPAddr) *EntryDataStu {
	var c EntryDataStu
	c.Init(Data, Addr)
	return &c
}

type EntryDataStu struct {
	Id      int
	NatType int
	Addr    *net.UDPAddr
	RData   []byte
	SData   []byte

	Kind int

	Role     int
	SipStart int
	Sport    int

	CliId int
}

func (r *EntryDataStu) Init(RData []byte, Addr *net.UDPAddr) (ret int) {
	r.RData = make([]byte, 128)
	r.SData = make([]byte, 128)

	r.RData = RData
	r.Addr = Addr
	if len(r.RData) < 8 {
		return
	}
	r.InitId()
	r.InitNatType()
	return
}
func (r *EntryDataStu) InitId() {
	tmp := r.RData[IdStart : IdStart+IdLen]
	r.Id = mylib.BytesToInt(tmp)
}

func (r *EntryDataStu) GetId() (Id string) {
	return strconv.Itoa(r.Id)
}
func (r *EntryDataStu) GetAddr() (Addr *net.UDPAddr) {
	return r.Addr
}

func (r *EntryDataStu) GetThis() interface{} {
	return r
}
func (r *EntryDataStu) SetId(Id int) {
	tmpByte := mylib.IntToBytes(Id)
	for i := 0; i < IdLen; i++ {
		r.SData[IdStart+i] = tmpByte[i]
	}
	return
}

func (r *EntryDataStu) SetCliId(Id int) {
	tmpByte := mylib.IntToBytes(Id)
	for i := 0; i < IdLen; i++ {
		r.SData[CliIdStart+i] = tmpByte[i]
	}
	return
}
func (r *EntryDataStu) GetCliId() (Id int) {
	tmpByte := r.RData[CliIdStart : CliIdStart+CliIdLen]
	return mylib.BytesToInt(tmpByte)
}
func (r *EntryDataStu) GetCliIdStr() (Id string) {
	return strconv.Itoa(r.GetCliId())
}

func (r *EntryDataStu) GetOtherAddr() (Addr *net.UDPAddr) {
	raddrstr := r.GetOtherIp() + ":" + r.GetOtherPort()
	raddr, err := net.ResolveUDPAddr("udp", raddrstr)
	if err != nil {
		fmt.Println("net.ResolveUDPAddr fail.", err)
		return
	}
	Addr = raddr
	return
}
func (r *EntryDataStu) GetOtherIp() (Ip string) {
	ip := net.IPv4(r.RData[OtherIpStart],
		r.RData[OtherIpStart+1],
		r.RData[OtherIpStart+2],
		r.RData[OtherIpStart+3])
	return ip.String()
}
func (r *EntryDataStu) GetOtherPort() (Port string) {
	port := int(r.RData[OtherPortStart])*256 + int(r.RData[OtherPortStart+1])
	return strconv.Itoa(port)
}
func (r *EntryDataStu) GetOtherPortInt() (Port int) {
	return int(r.RData[OtherPortStart])*256 + int(r.RData[OtherPortStart+1])
}
func (r *EntryDataStu) SetKind(Kind byte) {
	r.SData[KindStart] = Kind
	return
}
func (r *EntryDataStu) SetVedioStr(VData string) {
	for idx := 0; idx < len(VData); idx++ {
		r.SData[VedioDataStart+idx] = VData[idx]
	}
	return
}
func (r *EntryDataStu) GetVedioStr() (VData string) {
	return string(r.RData[VedioDataStart:])
}

func (r *EntryDataStu) SetRole(Role byte) {
	r.SData[RoleStart] = Role
	return
}
func (r *EntryDataStu) GetNatType() int {
	return r.NatType
}
func (r *EntryDataStu) GetNatTypeStr() string {
	return NatStr[r.NatType]
}
func (r *EntryDataStu) InitNatType() {
	r.NatType = int(r.RData[NatTypeStart])
}
func (r *EntryDataStu) SetNatType(NatType byte) {
	r.SData[NatTypeStart] = NatType
	return
}
func (r *EntryDataStu) GetIp() (Ip string) {
	ip := net.IPv4(r.RData[SipStart],
		r.RData[SipStart+1],
		r.RData[SipStart+2],
		r.RData[SipStart+3])
	return ip.String()
}
func (r *EntryDataStu) SetIp(IpStr string) {
	Ip := net.ParseIP(IpStr)
	ip4 := Ip.To4()
	for i := 0; i < SipLen; i++ {
		r.SData[SipStart+i] = ip4[i]
	}
	return
}

func (r *EntryDataStu) SetOtherAddr(Addr *net.UDPAddr) {
	Ip := Addr.IP
	ip4 := Ip.To4()
	for i := 0; i < SipLen; i++ {
		r.SData[OtherIpStart+i] = ip4[i]
	}
	r.SetOtherPort(uint16(Addr.Port))
	return
}
func (r *EntryDataStu) SetOtherPort(Port uint16) {
	r.SData[OtherPortStart] = byte(Port / 256)
	r.SData[OtherPortStart+1] = byte(Port % 256)
	return
}
func (r *EntryDataStu) SetPort(Port uint16) {
	r.SData[SportStart] = byte(Port / 256)
	r.SData[SportStart+1] = byte(Port % 256)
	return
}

func (r *EntryDataStu) String() (Rst string) {
	return string(r.SData)
}
func (r *EntryDataStu) Print() (Rst string) {
	Twlog := tool.NewTwLogStu("%v%t%v%t\n")
	defer Twlog.Flush()
	Twlog.Print("Id", r.Id)
	return
}
