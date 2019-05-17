package gonetworkmanager

//go:generate stringer -type=NmConnectivity
type NmConnectivity uint32

const (
	NmConnectivityUnknown NmConnectivity = 0
	NmConnectivityNone    NmConnectivity = 1
	NmConnectivityPortal  NmConnectivity = 2
	NmConnectivityLimited NmConnectivity = 3
	NmConnectivityFull    NmConnectivity = 4
)

//go:generate stringer -type=NmState
type NmState uint32

const (
	NmStateUnknown         NmState = 0
	NmStateAsleep          NmState = 10
	NmStateDisconnected    NmState = 20
	NmStateDisconnecting   NmState = 30
	NmStateConnecting      NmState = 40
	NmStateConnectedLocal  NmState = 50
	NmStateConnectedSite   NmState = 60
	NmStateConnectedGlobal NmState = 70
)

//go:generate stringer -type=NmDeviceState
type NmDeviceState uint32

const (
	NmDeviceStateUnknown      NmDeviceState = 0
	NmDeviceStateUnmanaged    NmDeviceState = 10
	NmDeviceStateUnavailable  NmDeviceState = 20
	NmDeviceStateDisconnected NmDeviceState = 30
	NmDeviceStatePrepare      NmDeviceState = 40
	NmDeviceStateConfig       NmDeviceState = 50
	NmDeviceStateNeed_auth    NmDeviceState = 60
	NmDeviceStateIp_config    NmDeviceState = 70
	NmDeviceStateIp_check     NmDeviceState = 80
	NmDeviceStateSecondaries  NmDeviceState = 90
	NmDeviceStateActivated    NmDeviceState = 100
	NmDeviceStateDeactivating NmDeviceState = 110
	NmDeviceStateFailed       NmDeviceState = 120
)

//go:generate stringer -type=NmActiveConnectionState
type NmActiveConnectionState uint32

const (
	NmActiveConnectionStateUnknown      NmActiveConnectionState = 0 // The state of the connection is unknown
	NmActiveConnectionStateActivating   NmActiveConnectionState = 1 // A network connection is being prepared
	NmActiveConnectionStateActivated    NmActiveConnectionState = 2 // There is a connection to the network
	NmActiveConnectionStateDeactivating NmActiveConnectionState = 3 // The network connection is being torn down and cleaned up
	NmActiveConnectionStateDeactivated  NmActiveConnectionState = 4 // The network connection is disconnected and will be removed
)

//go:generate stringer -type=NmDeviceType
type NmDeviceType uint32

const (
	NmDeviceTypeUnknown      NmDeviceType = 0  // unknown device
	NmDeviceTypeGeneric      NmDeviceType = 14 // generic support for unrecognized device types
	NmDeviceTypeEthernet     NmDeviceType = 1  // a wired ethernet device
	NmDeviceTypeWifi         NmDeviceType = 2  // an 802.11 Wi-Fi device
	NmDeviceTypeUnused1      NmDeviceType = 3  // not used
	NmDeviceTypeUnused2      NmDeviceType = 4  // not used
	NmDeviceTypeBt           NmDeviceType = 5  // a Bluetooth device supporting PAN or DUN access protocols
	NmDeviceTypeOlpcMesh     NmDeviceType = 6  // an OLPC XO mesh networking device
	NmDeviceTypeWimax        NmDeviceType = 7  // an 802.16e Mobile WiMAX broadband device
	NmDeviceTypeModem        NmDeviceType = 8  // a modem supporting analog telephone, CDMA/EVDO, GSM/UMTS, or LTE network access protocols
	NmDeviceTypeInfiniband   NmDeviceType = 9  // an IP-over-InfiniBand device
	NmDeviceTypeBond         NmDeviceType = 10 // a bond master interface
	NmDeviceTypeVlan         NmDeviceType = 11 // an 802.1Q VLAN interface
	NmDeviceTypeAdsl         NmDeviceType = 12 // ADSL modem
	NmDeviceTypeBridge       NmDeviceType = 13 // a bridge master interface
	NmDeviceTypeTeam         NmDeviceType = 15 // a team master interface
	NmDeviceTypeTun          NmDeviceType = 16 // a TUN or TAP interface
	NmDeviceTypeIpTunnel     NmDeviceType = 17 // a IP tunnel interface
	NmDeviceTypeMacvlan      NmDeviceType = 18 // a MACVLAN interface
	NmDeviceTypeVxlan        NmDeviceType = 19 // a VXLAN interface
	NmDeviceTypeVeth         NmDeviceType = 20 // a VETH interface
	NmDeviceTypeMacsec       NmDeviceType = 21 // a MACsec interface
	NmDeviceTypeDummy        NmDeviceType = 22 // a dummy interface
	NmDeviceTypePpp          NmDeviceType = 23 // a PPP interface
	NmDeviceTypeOvsInterface NmDeviceType = 24 // a Open vSwitch interface
	NmDeviceTypeOvsPort      NmDeviceType = 25 // a Open vSwitch port
	NmDeviceTypeOvsBridge    NmDeviceType = 26 // a Open vSwitch bridge
	NmDeviceTypeWpan         NmDeviceType = 27 // a IEEE 802.15.4 (WPAN) MAC Layer Device
	NmDeviceType6lowpan      NmDeviceType = 28 // 6LoWPAN interface
	NmDeviceTypeWireguard    NmDeviceType = 29 // a WireGuard interface
	NmDeviceTypeWifiP2p      NmDeviceType = 30 // an 802.11 Wi-Fi P2P device (Since: 1.16)
)

//go:generate stringer -type=Nm80211APFlags
type Nm80211APFlags uint32

const (
	Nm80211APFlagsNone    Nm80211APFlags = 0x0
	Nm80211APFlagsPrivacy Nm80211APFlags = 0x1
)

//go:generate stringer -type=Nm80211APSec
type Nm80211APSec uint32

const (
	Nm80211APSecNone         Nm80211APSec = 0x0
	Nm80211APSecPairWEP40    Nm80211APSec = 0x1
	Nm80211APSecPairWEP104   Nm80211APSec = 0x2
	Nm80211APSecPairTKIP     Nm80211APSec = 0x4
	Nm80211APSecPairCCMP     Nm80211APSec = 0x8
	Nm80211APSecGroupWEP40   Nm80211APSec = 0x10
	Nm80211APSecGroupWEP104  Nm80211APSec = 0x20
	Nm80211APSecGroupTKIP    Nm80211APSec = 0x40
	Nm80211APSecGroupCCMP    Nm80211APSec = 0x80
	Nm80211APSecKeyMgmtPSK   Nm80211APSec = 0x100
	Nm80211APSecKeyMgmt8021X Nm80211APSec = 0x200
)

//go:generate stringer -type=Nm80211Mode
type Nm80211Mode uint32

const (
	Nm80211ModeUnknown Nm80211Mode = 0
	Nm80211ModeAdhoc   Nm80211Mode = 1
	Nm80211ModeInfra   Nm80211Mode = 2
	Nm80211ModeAp      Nm80211Mode = 3
)
