package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus"
)

const (
	DeviceDummyInterface = DeviceInterface + ".Dummy"

	DeviceDummyPropertyHwAddress = DeviceDummyInterface + ".HwAddress" // readable   s
)

type DeviceDummy interface {
	Device

	// Hardware address of the device.
	GetHwAddress() string
}

func NewDeviceDummy(objectPath dbus.ObjectPath) (DeviceDummy, error) {
	var d deviceDummy
	return &d, d.init(NetworkManagerInterface, objectPath)
}

type deviceDummy struct {
	device
}

func (d *deviceDummy) GetHwAddress() string {
	return d.getStringProperty(DeviceDummyPropertyHwAddress)
}

func (d *deviceDummy) MarshalJSON() ([]byte, error) {
	m := d.device.marshalMap()
	m["HwAddress"] = d.GetHwAddress()
	return json.Marshal(m)
}
