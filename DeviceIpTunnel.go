package gonetworkmanager

import (
	"encoding/json"
	"github.com/godbus/dbus"
)

const (
	DeviceIpTunnelInterface = DeviceInterface + ".IPTunnel"

	DeviceIpTunnelPropertyHwAddress = DeviceIpTunnelInterface + "HwAddress" // readable   s

	DeviceIpTunnelPropertyMode               = DeviceIpTunnelInterface + ".Mode"               // readable   u
	DeviceIpTunnelPropertyParent             = DeviceIpTunnelInterface + ".Parent"             // readable   o
	DeviceIpTunnelPropertyLocal              = DeviceIpTunnelInterface + ".Local"              // readable   s
	DeviceIpTunnelPropertyRemote             = DeviceIpTunnelInterface + ".Remote"             // readable   s
	DeviceIpTunnelPropertyTtl                = DeviceIpTunnelInterface + ".Ttl"                // readable   y
	DeviceIpTunnelPropertyTos                = DeviceIpTunnelInterface + ".Tos"                // readable   y
	DeviceIpTunnelPropertyPathMtuDiscovery   = DeviceIpTunnelInterface + ".PathMtuDiscovery"   // readable   b
	DeviceIpTunnelPropertyInputKey           = DeviceIpTunnelInterface + ".InputKey"           // readable   s
	DeviceIpTunnelPropertyOutputKey          = DeviceIpTunnelInterface + ".OutputKey"          // readable   s
	DeviceIpTunnelPropertyEncapsulationLimit = DeviceIpTunnelInterface + ".EncapsulationLimit" // readable   y
	DeviceIpTunnelPropertyFlowLabel          = DeviceIpTunnelInterface + ".FlowLabel"          // readable   u
	DeviceIpTunnelPropertyFlags              = DeviceIpTunnelInterface + ".Flags"              // readable   u

)

type DeviceIpTunnel interface {
	Device

	// The tunneling mode
	GetMode() uint32

	// The object path of the parent device.
	GetParent() Device

	// The local endpoint of the tunnel.
	GetLocal() string

	// The remote endpoint of the tunnel.
	GetRemote() string

	// The TTL assigned to tunneled packets. 0 is a special value meaning that packets inherit the TTL value
	GetTtl() uint8

	// The type of service (IPv4) or traffic class (IPv6) assigned to tunneled packets.
	GetTos() uint8

	// Whether path MTU discovery is enabled on this tunnel.
	GetPathMtuDiscovery() bool

	// The key used for incoming packets.
	GetInputKey() string

	// The key used for outgoing packets.
	GetOutputKey() string

	// How many additional levels of encapsulation are permitted to be prepended to packets. This property applies only to IPv6 tunnels.
	GetEncapsulationLimit() uint8

	// The flow label to assign to tunnel packets. This property applies only to IPv6 tunnels.
	GetFlowLabel() uint32

	// Tunnel flags.
	GetFlags() uint32
}

func NewDeviceIpTunnel(objectPath dbus.ObjectPath) (DeviceIpTunnel, error) {
	var d deviceIpTunnel
	return &d, d.init(NetworkManagerInterface, objectPath)
}

type deviceIpTunnel struct {
	device
}

func (d *deviceIpTunnel) GetMode() uint32 {
	return d.getUint32Property(DeviceIpTunnelPropertyMode)
}

func (d *deviceIpTunnel) GetParent() Device {
	path := d.getObjectProperty(DeviceIpTunnelPropertyParent)
	if path == "/" {
		return nil
	}

	r, err := DeviceFactory(path)
	if err != nil {
		panic(err)
	}
	return r
}

func (d *deviceIpTunnel) GetLocal() string {
	return d.getStringProperty(DeviceIpTunnelPropertyLocal)
}

func (d *deviceIpTunnel) GetRemote() string {
	return d.getStringProperty(DeviceIpTunnelPropertyRemote)
}

func (d *deviceIpTunnel) GetTtl() uint8 {
	return d.getUint8Property(DeviceIpTunnelPropertyTtl)
}

func (d *deviceIpTunnel) GetTos() uint8 {
	return d.getUint8Property(DeviceIpTunnelPropertyTos)
}

func (d *deviceIpTunnel) GetPathMtuDiscovery() bool {
	return d.getBoolProperty(DeviceIpTunnelPropertyPathMtuDiscovery)
}

func (d *deviceIpTunnel) GetInputKey() string {
	return d.getStringProperty(DeviceIpTunnelPropertyInputKey)
}

func (d *deviceIpTunnel) GetOutputKey() string {
	return d.getStringProperty(DeviceIpTunnelPropertyOutputKey)
}

func (d *deviceIpTunnel) GetEncapsulationLimit() uint8 {
	return d.getUint8Property(DeviceIpTunnelPropertyEncapsulationLimit)
}

func (d *deviceIpTunnel) GetFlowLabel() uint32 {
	return d.getUint32Property(DeviceIpTunnelPropertyFlowLabel)
}

func (d *deviceIpTunnel) GetFlags() uint32 {
	return d.getUint32Property(DeviceIpTunnelPropertyFlags)
}

func (d *deviceIpTunnel) MarshalJSON() ([]uint8, error) {
	m := d.device.marshalMap()
	m["Mode"] = d.GetMode()
	m["Parent"] = d.GetParent()
	m["Local"] = d.GetLocal()
	m["Remote"] = d.GetRemote()
	m["Ttl"] = d.GetTtl()
	m["Tos"] = d.GetTos()
	m["PathMtuDiscovery"] = d.GetPathMtuDiscovery()
	m["InputKey"] = d.GetInputKey()
	m["OutputKey"] = d.GetOutputKey()
	m["EncapsulationLimit"] = d.GetEncapsulationLimit()
	m["FlowLabel"] = d.GetFlowLabel()
	m["Flags"] = d.GetFlags()
	return json.Marshal(m)
}
