package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus"
)

const (
	DeviceWiredInterface = DeviceInterface + ".Wired"

	DeviceWiredPropertyHwAddress       = DeviceWiredInterface + ".HwAddress"       // readable   s
	DeviceWiredPropertyPermHwAddress   = DeviceWiredInterface + ".PermHwAddress"   // readable   s
	DeviceWiredPropertySpeed           = DeviceWiredInterface + ".Speed"           // readable   u
	DeviceWiredPropertyS390Subchannels = DeviceWiredInterface + ".S390Subchannels" // readable   as
	DeviceWiredPropertyCarrier         = DeviceWiredInterface + ".Carrier"         // readable   b
)

type DeviceWired interface {
	Device

	// Active hardware address of the device.
	GetHwAddress() string

	// Permanent hardware address of the device.
	GetPermHwAddress() string

	// Design speed of the device, in megabits/second (Mb/s).
	GetSpeed() uint32

	// Array of S/390 subchannels for S/390 or z/Architecture devices.
	GetS390Subchannels() []string

	// Indicates whether the physical carrier is found (e.g. whether a cable is plugged in or not).
	GetCarrier() bool
}

func NewDeviceWired(objectPath dbus.ObjectPath) (DeviceWired, error) {
	var d deviceWired
	return &d, d.init(NetworkManagerInterface, objectPath)
}

type deviceWired struct {
	device
}

func (d *deviceWired) GetHwAddress() string {
	return d.getStringProperty(DeviceWiredPropertyHwAddress)
}

func (d *deviceWired) GetPermHwAddress() string {
	return d.getStringProperty(DeviceWiredPropertyPermHwAddress)
}

func (d *deviceWired) GetSpeed() uint32 {
	return d.getUint32Property(DeviceWiredPropertySpeed)
}

func (d *deviceWired) GetS390Subchannels() []string {
	return d.getSliceStringProperty(DeviceWiredPropertyS390Subchannels)
}

func (d *deviceWired) GetCarrier() bool {
	return d.getBoolProperty(DeviceWiredPropertyCarrier)
}

func (d *deviceWired) MarshalJSON() ([]byte, error) {
	m := d.device.marshalMap()
	m["HwAddress"] = d.GetHwAddress()
	m["PermHwAddress"] = d.GetPermHwAddress()
	m["Speed"] = d.GetSpeed()
	m["S390Subchannels"] = d.GetS390Subchannels()
	m["Carrier"] = d.GetCarrier()
	return json.Marshal(m)
}
