package base

var KindStart int = 0
var KindLen int = 1

var NatTypeStart int = 1
var NatTypeLen int = 1

var RoleStart int = 2
var RoleLen int = 1

var SipStart int = 3
var SipLen int = 4

var IdStart int = 4
var IdLen int = 4

var SportStart int = 12
var SportLen int = 2

var CliIdStart int = 20
var CliIdLen int = 4

var OtherIpStart int = 20
var OtherIpLen int = 4

var OtherPortStart int = 24
var OtherPortLen int = 2

var VedioDataStart int = 20
var VedioDatalen int = 500

var EntryKindHeart int = 110

const (
	NATError int = iota
	NATUnknown
	NATNone
	NATBlocked
	NATFull
	NATSymmetric
	NATRestricted
	NATPortRestricted
	NATSymmetricUDPFirewall

	NATSymetric            = NATSymmetric
	NATSymetricUDPFirewall = NATSymmetricUDPFirewall
)

var NatStr map[int]string

func init() {
	NatStr = map[int]string{
		NATError:                "Test failed",
		NATUnknown:              "Unexpected response from the STUN server",
		NATBlocked:              "UDP is blocked",
		NATFull:                 "KindFull",
		NATSymmetric:            "KindSymmetric",
		NATRestricted:           "KindRestricted",
		NATPortRestricted:       "KindPortrestricted",
		NATNone:                 "Not behind a NAT",
		NATSymmetricUDPFirewall: "KindSymmetricUDPFirewall",
	}
}
