package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus"
)

const (
	DeviceInterface = NetworkManagerInterface + ".Device"

	/* Properties */
	DevicePropertyUdi                  = DeviceInterface + ".Udi"                  // readable   s
	DevicePropertyInterface            = DeviceInterface + ".Interface"            // readable   s
	DevicePropertyIpInterface          = DeviceInterface + ".IpInterface"          // readable   s
	DevicePropertyDriver               = DeviceInterface + ".Driver"               // readable   s
	DevicePropertyDriverVersion        = DeviceInterface + ".DriverVersion"        // readable   s
	DevicePropertyFirmwareVersion      = DeviceInterface + ".FirmwareVersion"      // readable   s
	DevicePropertyCapabilities         = DeviceInterface + ".Capabilities"         // readable   u
	DevicePropertyIp4Address           = DeviceInterface + ".Ip4Address"           // readable   u
	DevicePropertyState                = DeviceInterface + ".State"                // readable   u
	DevicePropertyStateReason          = DeviceInterface + ".StateReason"          // readable   (uu)
	DevicePropertyActiveConnection     = DeviceInterface + ".ActiveConnection"     // readable   o
	DevicePropertyIp4Config            = DeviceInterface + ".Ip4Config"            // readable   o
	DevicePropertyDhcp4Config          = DeviceInterface + ".Dhcp4Config"          // readable   o
	DevicePropertyIp6Config            = DeviceInterface + ".Ip6Config"            // readable   o
	DevicePropertyDhcp6Config          = DeviceInterface + ".Dhcp6Config"          // readable   o
	DevicePropertyManaged              = DeviceInterface + ".Managed"              // readwrite  b
	DevicePropertyAutoconnect          = DeviceInterface + ".Autoconnect"          // readwrite  b
	DevicePropertyFirmwareMissing      = DeviceInterface + ".FirmwareMissing"      // readable   b
	DevicePropertyNmPluginMissing      = DeviceInterface + ".NmPluginMissing"      // readable   b
	DevicePropertyDeviceType           = DeviceInterface + ".DeviceType"           // readable   u
	DevicePropertyAvailableConnections = DeviceInterface + ".AvailableConnections" // readable   ao
	DevicePropertyPhysicalPortId       = DeviceInterface + ".PhysicalPortId"       // readable   s
	DevicePropertyMtu                  = DeviceInterface + ".Mtu"                  // readable   u
	DevicePropertyMetered              = DeviceInterface + ".Metered"              // readable   u
	DevicePropertyLldpNeighbors        = DeviceInterface + ".LldpNeighbors"        // readable   aa{sv}
	DevicePropertyReal                 = DeviceInterface + ".Real"                 // readable   b
	DevicePropertyIp4Connectivity      = DeviceInterface + ".Ip4Connectivity"      // readable   u
)

func DeviceFactory(objectPath dbus.ObjectPath) (Device, error) {
	d, err := NewDevice(objectPath)
	if err != nil {
		return nil, err
	}

	switch d.GetDeviceType() {
	case NmDeviceTypeWifi:
		return NewWirelessDevice(objectPath)
	}

	return d, nil
}

type Device interface {
	GetPath() dbus.ObjectPath

	// Operating-system specific transient device hardware identifier. This is an
	// opaque string representing the underlying hardware for the device, and
	// shouldn't be used to keep track of individual devices. For some device types
	// (Bluetooth, Modems) it is an identifier used by the hardware service
	// (ie bluez or ModemManager) to refer to that device, and client programs use
	// it get additional information from those services which NM does not provide.
	// The Udi is not guaranteed to be consistent across reboots or hotplugs of the
	// hardware. If you're looking for a way to uniquely track each device in your
	// application, use the object path. If you're looking for a way to track a
	// specific piece of hardware across reboot or hotplug, use a MAC address or
	// USB serial number.
	GetUdi() string

	// GetInterface gets the name of the device's control (and often data)
	// interface.
	GetInterface() string

	// GetIpInterface gets the IP interface name of the device.
	GetIpInterface() string

	// GetState gets the current state of the device.
	GetState() NmDeviceState

	// GetIP4Config gets the Ip4Config object describing the configuration of the
	// device. Only valid when the device is in the NM_DEVICE_STATE_ACTIVATED
	// state.
	GetIP4Config() IP4Config

	// GetDHCP4Config gets the Dhcp4Config object describing the configuration of the
	// device. Only valid when the device is in the NM_DEVICE_STATE_ACTIVATED
	// state.
	GetDHCP4Config() DHCP4Config

	// GetDeviceType gets the general type of the network device; ie Ethernet,
	// WiFi, etc.
	GetDeviceType() NmDeviceType

	// GetAvailableConnections gets an array of object paths of every configured
	// connection that is currently 'available' through this device.
	GetAvailableConnections() []Connection

	// Whether or not this device is managed by NetworkManager. Setting this property
	// has a similar effect to configuring the device as unmanaged via the
	// keyfile.unmanaged-devices setting in NetworkManager.conf. Changes to this value
	// are not persistent and lost after NetworkManager restart.
	GetManaged() bool

	// If TRUE, indicates the device is allowed to autoconnect. If FALSE, manual intervention is required before the device will automatically connect to a known network, such as activating a connection using the device, or setting this property to TRUE. This property cannot be set to TRUE for default-unmanaged devices, since they never autoconnect.
	GetAutoConnect() bool

	MarshalJSON() ([]byte, error)
}

func NewDevice(objectPath dbus.ObjectPath) (Device, error) {
	var d device
	return &d, d.init(NetworkManagerInterface, objectPath)
}

type device struct {
	dbusBase
}

func (d *device) GetPath() dbus.ObjectPath {
	return d.obj.Path()
}

func (d *device) GetUdi() string {
	return d.getStringProperty(DevicePropertyUdi)
}

func (d *device) GetInterface() string {
	return d.getStringProperty(DevicePropertyInterface)
}

func (d *device) GetIpInterface() string {
	return d.getStringProperty(DevicePropertyIpInterface)
}

func (d *device) GetState() NmDeviceState {
	return NmDeviceState(d.getUint32Property(DevicePropertyState))
}

func (d *device) GetIP4Config() IP4Config {
	path := d.getObjectProperty(DevicePropertyIp4Config)
	if path == "/" {
		return nil
	}

	cfg, err := NewIP4Config(path)
	if err != nil {
		panic(err)
	}

	return cfg
}

func (d *device) GetDHCP4Config() DHCP4Config {
	path := d.getObjectProperty(DevicePropertyDhcp4Config)
	if path == "/" {
		return nil
	}

	cfg, err := NewDHCP4Config(path)
	if err != nil {
		panic(err)
	}

	return cfg
}

func (d *device) GetDeviceType() NmDeviceType {
	return NmDeviceType(d.getUint32Property(DevicePropertyDeviceType))
}

func (d *device) GetAvailableConnections() []Connection {
	connPaths := d.getSliceObjectProperty(DevicePropertyAvailableConnections)
	conns := make([]Connection, len(connPaths))

	var err error
	for i, path := range connPaths {
		conns[i], err = NewConnection(path)
		if err != nil {
			panic(err)
		}
	}

	return conns
}

func (d *device) GetManaged() bool {
	return d.getBoolProperty(DevicePropertyManaged)
}

func (d *device) GetAutoConnect() bool {
	return d.getBoolProperty(DevicePropertyAutoconnect)
}

func (d *device) marshalMap() map[string]interface{} {
	return map[string]interface{}{
		"Interface":            d.GetInterface(),
		"IP interface":         d.GetIpInterface(),
		"State":                d.GetState().String(),
		"IP4Config":            d.GetIP4Config(),
		"DHCP4Config":          d.GetDHCP4Config(),
		"DeviceType":           d.GetDeviceType().String(),
		"AvailableConnections": d.GetAvailableConnections(),
	}
}

func (d *device) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.marshalMap())
}
