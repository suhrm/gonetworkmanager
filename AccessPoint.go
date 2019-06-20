package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus"
)

const (
	AccessPointInterface = NetworkManagerInterface + ".AccessPoint"

	/* Properties */
	AccessPointPropertyFlags      = AccessPointInterface + ".Flags"      // readable   u
	AccessPointPropertyWpaFlags   = AccessPointInterface + ".WpaFlags"   // readable   u
	AccessPointPropertyRsnFlags   = AccessPointInterface + ".RsnFlags"   // readable   u
	AccessPointPropertySsid       = AccessPointInterface + ".Ssid"       // readable   ay
	AccessPointPropertyFrequency  = AccessPointInterface + ".Frequency"  // readable   u
	AccessPointPropertyHwAddress  = AccessPointInterface + ".HwAddress"  // readable   s
	AccessPointPropertyMode       = AccessPointInterface + ".Mode"       // readable   u
	AccessPointPropertyMaxBitrate = AccessPointInterface + ".MaxBitrate" // readable   u
	AccessPointPropertyStrength   = AccessPointInterface + ".Strength"   // readable   y
	AccessPointPropertyLastSeen   = AccessPointInterface + ".LastSeen"   // readable   i
)

type AccessPoint interface {
	GetPath() dbus.ObjectPath

	// GetFlags gets flags describing the capabilities of the access point.
	GetPropertyFlags() (uint32, error)

	// GetWPAFlags gets flags describing the access point's capabilities
	// according to WPA (Wifi Protected Access).
	GetPropertyWPAFlags() (uint32, error)

	// GetRSNFlags gets flags describing the access point's capabilities
	// according to the RSN (Robust Secure Network) protocol.
	GetPropertyRSNFlags() (uint32, error)

	// GetSSID returns the Service Set Identifier identifying the access point.
	GetPropertySSID() (string, error)

	// GetFrequency gets the radio channel frequency in use by the access point,
	// in MHz.
	GetPropertyFrequency() (uint32, error)

	// GetHWAddress gets the hardware address (BSSID) of the access point.
	GetPropertyHWAddress() (string, error)

	// GetMode describes the operating mode of the access point.
	GetPropertyMode() (Nm80211Mode, error)

	// GetMaxBitrate gets the maximum bitrate this access point is capable of, in
	// kilobits/second (Kb/s).
	GetPropertyMaxBitrate() (uint32, error)

	// GetStrength gets the current signal quality of the access point, in
	// percent.
	GetPropertyStrength() (uint8, error)

	MarshalJSON() ([]byte, error)
}

func NewAccessPoint(objectPath dbus.ObjectPath) (AccessPoint, error) {
	var a accessPoint
	return &a, a.init(NetworkManagerInterface, objectPath)
}

type accessPoint struct {
	dbusBase
}

func (a *accessPoint) GetPath() dbus.ObjectPath {
	return a.obj.Path()
}

func (a *accessPoint) GetPropertyFlags() (uint32, error) {
	return a.getUint32Property(AccessPointPropertyFlags)
}

func (a *accessPoint) GetPropertyWPAFlags() (uint32, error) {
	return a.getUint32Property(AccessPointPropertyWpaFlags)
}

func (a *accessPoint) GetPropertyRSNFlags() (uint32, error) {
	return a.getUint32Property(AccessPointPropertyRsnFlags)
}

func (a *accessPoint) GetPropertySSID() (string, error) {
	v, err := a.getSliceByteProperty(AccessPointPropertySsid)
	return string(v), err
}

func (a *accessPoint) GetPropertyFrequency() (uint32, error) {
	return a.getUint32Property(AccessPointPropertyFrequency)
}

func (a *accessPoint) GetPropertyHWAddress() (string, error) {
	return a.getStringProperty(AccessPointPropertyHwAddress)
}

func (a *accessPoint) GetPropertyMode() (Nm80211Mode, error) {
	v, err := a.getUint32Property(AccessPointPropertyMode)
	return Nm80211Mode(v), err
}

func (a *accessPoint) GetPropertyMaxBitrate() (uint32, error) {
	return a.getUint32Property(AccessPointPropertyMaxBitrate)
}

func (a *accessPoint) GetPropertyStrength() (uint8, error) {
	return a.getUint8Property(AccessPointPropertyStrength)
}

func (a *accessPoint) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})

	m["Flags"], _ = a.GetPropertyFlags()
	m["WPAFlags"], _ = a.GetPropertyWPAFlags()
	m["RSNFlags"], _ = a.GetPropertyRSNFlags()
	m["SSID"], _ = a.GetPropertySSID()
	m["Frequency"], _ = a.GetPropertyFrequency()
	m["HWAddress"], _ = a.GetPropertyHWAddress()
	m["MaxBitrate"], _ = a.GetPropertyMaxBitrate()
	m["Strength"], _ = a.GetPropertyStrength()

	mode, _ := a.GetPropertyMode()
	m["Mode"] = mode.String()

	return json.Marshal(m)
}
