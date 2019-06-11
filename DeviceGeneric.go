package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus"
)

const (
	DeviceGenericInterface = DeviceInterface + ".Generic"

	DeviceGenericPropertyHwAddress       = DeviceGenericInterface + ".HwAddress"       // readable   s
	DeviceGenericPropertyTypeDescription = DeviceGenericInterface + ".TypeDescription" // readable   s
)

type DeviceGeneric interface {
	Device

	// Active hardware address of the device.
	GetHwAddress() string

	// A (non-localized) description of the interface type, if known.
	GetTypeDescription() string
}

func NewDeviceGeneric(objectPath dbus.ObjectPath) (DeviceGeneric, error) {
	var d deviceGeneric
	return &d, d.init(NetworkManagerInterface, objectPath)
}

type deviceGeneric struct {
	device
}

func (d *deviceGeneric) GetHwAddress() string {
	return d.getStringProperty(DeviceGenericPropertyHwAddress)
}

func (d *deviceGeneric) GetTypeDescription() string {
	return d.getStringProperty(DeviceGenericPropertyTypeDescription)
}

func (d *deviceGeneric) MarshalJSON() ([]byte, error) {
	m := d.device.marshalMap()
	m["HwAddress"] = d.GetHwAddress()
	m["TypeDescription"] = d.GetTypeDescription()
	return json.Marshal(m)
}
