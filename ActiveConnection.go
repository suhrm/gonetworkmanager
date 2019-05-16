package gonetworkmanager

import (
	"github.com/godbus/dbus"
)

const (
	ActiveConnectionInterface = NetworkManagerInterface + ".Connection.Active"

	/* Property */
	ActiveConnectionProperyConnection     = ActiveConnectionInterface + ".Connection"     // readable   o
	ActiveConnectionProperySpecificObject = ActiveConnectionInterface + ".SpecificObject" // readable   o
	ActiveConnectionProperyId             = ActiveConnectionInterface + ".Id"             // readable   s
	ActiveConnectionProperyUuid           = ActiveConnectionInterface + ".Uuid"           // readable   s
	ActiveConnectionProperyType           = ActiveConnectionInterface + ".Type"           // readable   s
	ActiveConnectionProperyDevices        = ActiveConnectionInterface + ".Devices"        // readable   ao
	ActiveConnectionProperyState          = ActiveConnectionInterface + ".State"          // readable   u
	ActiveConnectionProperyStateFlags     = ActiveConnectionInterface + ".StateFlags"     // readable   u
	ActiveConnectionProperyDefault        = ActiveConnectionInterface + ".Default"        // readable   b
	ActiveConnectionProperyIp4Config      = ActiveConnectionInterface + ".Ip4Config"      // readable   o
	ActiveConnectionProperyDhcp4Config    = ActiveConnectionInterface + ".Dhcp4Config"    // readable   o
	ActiveConnectionProperyDefault6       = ActiveConnectionInterface + ".Default6"       // readable   b
	ActiveConnectionProperyIp6Config      = ActiveConnectionInterface + ".Ip6Config"      // readable   o
	ActiveConnectionProperyDhcp6Config    = ActiveConnectionInterface + ".Dhcp6Config"    // readable   o
	ActiveConnectionProperyVpn            = ActiveConnectionInterface + ".Vpn"            // readable   b
	ActiveConnectionProperyMaster         = ActiveConnectionInterface + ".Master"         // readable   o
)

type ActiveConnection interface {
	// GetConnection gets connection object of the connection.
	GetConnection() Connection

	// GetSpecificObject gets a specific object associated with the active connection.
	GetSpecificObject() AccessPoint

	// GetID gets the ID of the connection.
	GetID() string

	// GetUUID gets the UUID of the connection.
	GetUUID() string

	// GetType gets the type of the connection.
	GetType() string

	// GetDevices gets array of device objects which are part of this active connection.
	GetDevices() []Device

	// GetState gets the state of the connection.
	GetState() uint32

	// GetStateFlags gets the state flags of the connection.
	GetStateFlags() uint32

	// GetDefault gets the default IPv4 flag of the connection.
	GetDefault() bool

	// GetIP4Config gets the IP4Config of the connection.
	GetIP4Config() IP4Config

	// GetDHCP4Config gets the DHCP4Config of the connection.
	GetDHCP4Config() DHCP4Config

	// GetVPN gets the VPN flag of the connection.
	GetVPN() bool

	// GetMaster gets the master device of the connection.
	GetMaster() Device
}

func NewActiveConnection(objectPath dbus.ObjectPath) (ActiveConnection, error) {
	var a activeConnection
	return &a, a.init(NetworkManagerInterface, objectPath)
}

type activeConnection struct {
	dbusBase
}

func (a *activeConnection) GetConnection() Connection {
	path := a.getObjectProperty(ActiveConnectionProperyConnection)
	con, err := NewConnection(path)
	if err != nil {
		panic(err)
	}
	return con
}

func (a *activeConnection) GetSpecificObject() AccessPoint {
	path := a.getObjectProperty(ActiveConnectionProperySpecificObject)
	ap, err := NewAccessPoint(path)
	if err != nil {
		panic(err)
	}
	return ap
}

func (a *activeConnection) GetID() string {
	return a.getStringProperty(ActiveConnectionProperyId)
}

func (a *activeConnection) GetUUID() string {
	return a.getStringProperty(ActiveConnectionProperyUuid)
}

func (a *activeConnection) GetType() string {
	return a.getStringProperty(ActiveConnectionProperyType)
}

func (a *activeConnection) GetDevices() []Device {
	paths := a.getSliceObjectProperty(ActiveConnectionProperyDevices)
	devices := make([]Device, len(paths))
	var err error
	for i, path := range paths {
		devices[i], err = DeviceFactory(path)
		if err != nil {
			panic(err)
		}
	}
	return devices
}

func (a *activeConnection) GetState() uint32 {
	return a.getUint32Property(ActiveConnectionProperyState)
}

func (a *activeConnection) GetStateFlags() uint32 {
	return a.getUint32Property(ActiveConnectionProperyStateFlags)
}

func (a *activeConnection) GetDefault() bool {
	b := a.getProperty(ActiveConnectionProperyDefault)
	return b.(bool)
}

func (a *activeConnection) GetIP4Config() IP4Config {
	path := a.getObjectProperty(ActiveConnectionProperyIp4Config)
	r, err := NewIP4Config(path)
	if err != nil {
		panic(err)
	}
	return r
}

func (a *activeConnection) GetDHCP4Config() DHCP4Config {
	path := a.getObjectProperty(ActiveConnectionProperyDhcp4Config)
	r, err := NewDHCP4Config(path)
	if err != nil {
		panic(err)
	}
	return r
}

func (a *activeConnection) GetVPN() bool {
	ret := a.getProperty(ActiveConnectionProperyVpn)
	return ret.(bool)
}

func (a *activeConnection) GetMaster() Device {
	path := a.getObjectProperty(ActiveConnectionProperyMaster)
	r, err := DeviceFactory(path)
	if err != nil {
		panic(err)
	}
	return r
}
