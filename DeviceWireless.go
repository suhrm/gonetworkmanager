package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus"
)

const (
	DeviceWirelessInterface = DeviceInterface + ".Wireless"

	// Methods
	DeviceWirelessGetAccessPoints = DeviceWirelessInterface + ".GetAccessPoints"
	DeviceWirelessRequestScan     = DeviceWirelessInterface + ".RequestScan"
)

type DeviceWireless interface {
	Device

	// GetAccessPoints gets the list of access points visible to this device.
	// Note that this list does not include access points which hide their SSID.
	// To retrieve a list of all access points (including hidden ones) use the
	// GetAllAccessPoints() method.
	GetAccessPoints() ([]AccessPoint, error)

	RequestScan()
}

func NewDeviceWireless(objectPath dbus.ObjectPath) (DeviceWireless, error) {
	var d deviceWireless
	return &d, d.init(NetworkManagerInterface, objectPath)
}

type deviceWireless struct {
	device
}

func (d *deviceWireless) GetAccessPoints() ([]AccessPoint, error) {
	var apPaths []dbus.ObjectPath
	err := d.callWithReturn(&apPaths, DeviceWirelessGetAccessPoints)

	if err != nil {
		return nil, err
	}

	aps := make([]AccessPoint, len(apPaths))

	for i, path := range apPaths {
		aps[i], err = NewAccessPoint(path)
		if err != nil {
			return aps, err
		}
	}

	return aps, nil
}

func (d *deviceWireless) RequestScan() {
	var options map[string]interface{}
	d.obj.Call(DeviceWirelessRequestScan, 0, options).Store()
}

func (d *deviceWireless) MarshalJSON() ([]byte, error) {
	m := d.device.marshalMap()
	m["AccessPoints"], _ = d.GetAccessPoints()
	return json.Marshal(m)
}
