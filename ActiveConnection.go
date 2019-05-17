package gonetworkmanager

import (
	"github.com/godbus/dbus"
)

const (
	ActiveConnectionInterface = NetworkManagerInterface + ".Connection.Active"

	/* Property */
	ActiveConnectionPropertyConnection     = ActiveConnectionInterface + ".Connection"     // readable   o
	ActiveConnectionPropertySpecificObject = ActiveConnectionInterface + ".SpecificObject" // readable   o
	ActiveConnectionPropertyId             = ActiveConnectionInterface + ".Id"             // readable   s
	ActiveConnectionPropertyUuid           = ActiveConnectionInterface + ".Uuid"           // readable   s
	ActiveConnectionPropertyType           = ActiveConnectionInterface + ".Type"           // readable   s
	ActiveConnectionPropertyDevices        = ActiveConnectionInterface + ".Devices"        // readable   ao
	ActiveConnectionPropertyState          = ActiveConnectionInterface + ".State"          // readable   u
	ActiveConnectionPropertyStateFlags     = ActiveConnectionInterface + ".StateFlags"     // readable   u
	ActiveConnectionPropertyDefault        = ActiveConnectionInterface + ".Default"        // readable   b
	ActiveConnectionPropertyIp4Config      = ActiveConnectionInterface + ".Ip4Config"      // readable   o
	ActiveConnectionPropertyDhcp4Config    = ActiveConnectionInterface + ".Dhcp4Config"    // readable   o
	ActiveConnectionPropertyDefault6       = ActiveConnectionInterface + ".Default6"       // readable   b
	ActiveConnectionPropertyIp6Config      = ActiveConnectionInterface + ".Ip6Config"      // readable   o
	ActiveConnectionPropertyDhcp6Config    = ActiveConnectionInterface + ".Dhcp6Config"    // readable   o
	ActiveConnectionPropertyVpn            = ActiveConnectionInterface + ".Vpn"            // readable   b
	ActiveConnectionPropertyMaster         = ActiveConnectionInterface + ".Master"         // readable   o
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
	GetState() NmActiveConnectionState

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
	path := a.getObjectProperty(ActiveConnectionPropertyConnection)
	con, err := NewConnection(path)
	if err != nil {
		panic(err)
	}
	return con
}

func (a *activeConnection) GetSpecificObject() AccessPoint {
	path := a.getObjectProperty(ActiveConnectionPropertySpecificObject)
	ap, err := NewAccessPoint(path)
	if err != nil {
		panic(err)
	}
	return ap
}

func (a *activeConnection) GetID() string {
	return a.getStringProperty(ActiveConnectionPropertyId)
}

func (a *activeConnection) GetUUID() string {
	return a.getStringProperty(ActiveConnectionPropertyUuid)
}

func (a *activeConnection) GetType() string {
	return a.getStringProperty(ActiveConnectionPropertyType)
}

func (a *activeConnection) GetDevices() []Device {
	paths := a.getSliceObjectProperty(ActiveConnectionPropertyDevices)
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

func (a *activeConnection) GetState() NmActiveConnectionState {
	return NmActiveConnectionState(a.getUint32Property(ActiveConnectionPropertyState))
}

func (a *activeConnection) GetStateFlags() uint32 {
	return a.getUint32Property(ActiveConnectionPropertyStateFlags)
}

func (a *activeConnection) GetDefault() bool {
	b := a.getProperty(ActiveConnectionPropertyDefault)
	return b.(bool)
}

func (a *activeConnection) GetIP4Config() IP4Config {
	path := a.getObjectProperty(ActiveConnectionPropertyIp4Config)
	r, err := NewIP4Config(path)
	if err != nil {
		panic(err)
	}
	return r
}

func (a *activeConnection) GetDHCP4Config() DHCP4Config {
	path := a.getObjectProperty(ActiveConnectionPropertyDhcp4Config)
	r, err := NewDHCP4Config(path)
	if err != nil {
		panic(err)
	}
	return r
}

func (a *activeConnection) GetVPN() bool {
	ret := a.getProperty(ActiveConnectionPropertyVpn)
	return ret.(bool)
}

func (a *activeConnection) GetMaster() Device {
	path := a.getObjectProperty(ActiveConnectionPropertyMaster)
	r, err := DeviceFactory(path)
	if err != nil {
		panic(err)
	}
	return r
}
