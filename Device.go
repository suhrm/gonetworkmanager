package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus"
)

const (
	DeviceInterface = NetworkManagerInterface + ".Device"

	/* Methods */
	DeviceReapply              = DeviceInterface + ".Reapply"
	DeviceGetAppliedConnection = DeviceInterface + ".GetAppliedConnection"
	DeviceDisconnect           = DeviceInterface + ".Disconnect"
	DeviceDelete               = DeviceInterface + ".Delete"

	/* Properties */
	DevicePropertyUdi                  = DeviceInterface + ".Udi"                  // readable   s
	DevicePropertyInterface            = DeviceInterface + ".Interface"            // readable   s
	DevicePropertyIpInterface          = DeviceInterface + ".IpInterface"          // readable   s
	DevicePropertyDriver               = DeviceInterface + ".Driver"               // readable   s
	DevicePropertyDriverVersion        = DeviceInterface + ".DriverVersion"        // readable   s
	DevicePropertyFirmwareVersion      = DeviceInterface + ".FirmwareVersion"      // readable   s
	DevicePropertyCapabilities         = DeviceInterface + ".Capabilities"         // readable   u
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

	deviceType, err := d.GetPropertyDeviceType()
	if err != nil {
		return nil, err
	}

	switch deviceType {
	case NmDeviceTypeDummy:
		return NewDeviceDummy(objectPath)
	case NmDeviceTypeGeneric:
		return NewDeviceGeneric(objectPath)
	case NmDeviceTypeIpTunnel:
		return NewDeviceIpTunnel(objectPath)
	case NmDeviceTypeEthernet:
		return NewDeviceWired(objectPath)
	case NmDeviceTypeWifi:
		return NewDeviceWireless(objectPath)
	}

	return d, nil
}

type Device interface {
	GetPath() dbus.ObjectPath

	// Disconnects a device and prevents the device from automatically activating further connections without user intervention.
	Disconnect() error

	// Deletes a software device from NetworkManager and removes the interface from the system. The method returns an error when called for a hardware device.
	Delete() error

	// Operating-system specific transient device hardware identifier. This is an opaque string representing the underlying hardware for the device, and shouldn't be used to keep track of individual devices. For some device types (Bluetooth, Modems) it is an identifier used by the hardware service (ie bluez or ModemManager) to refer to that device, and client programs use it get additional information from those services which NM does not provide. The Udi is not guaranteed to be consistent across reboots or hotplugs of the hardware. If you're looking for a way to uniquely track each device in your application, use the object path. If you're looking for a way to track a specific piece of hardware across reboot or hotplug, use a MAC address or USB serial number.
	GetPropertyUdi() (string, error)

	// The name of the device's control (and often data) interface. Note that non UTF-8 characters are backslash escaped, so the resulting name may be longer then 15 characters. Use g_strcompress() to revert the escaping.
	GetPropertyInterface() (string, error)

	// The name of the device's data interface when available. This property may not refer to the actual data interface until the device has successfully established a data connection, indicated by the device's State becoming ACTIVATED. Note that non UTF-8 characters are backslash escaped, so the resulting name may be longer then 15 characters. Use g_strcompress() to revert the escaping.
	GetPropertyIpInterface() (string, error)

	// The driver handling the device. Non-UTF-8 sequences are backslash escaped. Use g_strcompress() to revert.
	GetPropertyDriver() (string, error)

	// The version of the driver handling the device. Non-UTF-8 sequences are backslash escaped. Use g_strcompress() to revert.
	GetPropertyDriverVersion() (string, error)

	// The firmware version for the device. Non-UTF-8 sequences are backslash escaped. Use g_strcompress() to revert.
	GetPropertyFirmwareVersion() (string, error)

	// The current state of the device.
	GetPropertyState() (NmDeviceState, error)

	// Object path of the Ip4Config object describing the configuration of the device. Only valid when the device is in the NM_DEVICE_STATE_ACTIVATED state.
	GetPropertyIP4Config() (IP4Config, error)

	// Object path of the Dhcp4Config object describing the DHCP options returned by the DHCP server. Only valid when the device is in the NM_DEVICE_STATE_ACTIVATED state.
	GetPropertyDHCP4Config() (DHCP4Config, error)

	// Object path of the Ip6Config object describing the configuration of the device. Only valid when the device is in the NM_DEVICE_STATE_ACTIVATED state.
	GetPropertyIP6Config() (IP6Config, error)

	// Object path of the Dhcp6Config object describing the DHCP options returned by the DHCP server. Only valid when the device is in the NM_DEVICE_STATE_ACTIVATED state.
	GetPropertyDHCP6Config() (DHCP6Config, error)

	// Whether or not this device is managed by NetworkManager. Setting this property has a similar effect to configuring the device as unmanaged via the keyfile.unmanaged-devices setting in NetworkManager.conf. Changes to this value are not persistent and lost after NetworkManager restart.
	GetPropertyManaged() (bool, error)

	// If TRUE, indicates the device is allowed to autoconnect. If FALSE, manual intervention is required before the device will automatically connect to a known network, such as activating a connection using the device, or setting this property to TRUE. This property cannot be set to TRUE for default-unmanaged devices, since they never autoconnect.
	GetPropertyAutoConnect() (bool, error)

	// If TRUE, indicates the device is likely missing firmware necessary for its operation.
	GetPropertyFirmwareMissing() (bool, error)

	// If TRUE, indicates the NetworkManager plugin for the device is likely missing or misconfigured.
	GetPropertyNmPluginMissing() (bool, error)

	// The general type of the network device; ie Ethernet, Wi-Fi, etc.
	GetPropertyDeviceType() (NmDeviceType, error)

	// An array of object paths of every configured connection that is currently 'available' through this device.
	GetPropertyAvailableConnections() ([]Connection, error)

	// If non-empty, an (opaque) indicator of the physical network port associated with the device. This can be used to recognize when two seemingly-separate hardware devices are actually just different virtual interfaces to the same physical port.
	GetPropertyPhysicalPortId() (string, error)

	// The device MTU (maximum transmission unit).
	GetPropertyMtu() (uint32, error)

	// True if the device exists, or False for placeholder devices that do not yet exist but could be automatically created by NetworkManager if one of their AvailableConnections was activated.
	GetPropertyReal() (bool, error)

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

func (d *device) Disconnect() error {
	return d.call(DeviceDisconnect)
}

func (d *device) Delete() error {
	return d.call(DeviceDelete)
}

func (d *device) GetPropertyUdi() (string, error) {
	return d.getStringProperty(DevicePropertyUdi)
}

func (d *device) GetPropertyInterface() (string, error) {
	return d.getStringProperty(DevicePropertyInterface)
}

func (d *device) GetPropertyIpInterface() (string, error) {
	return d.getStringProperty(DevicePropertyIpInterface)
}

func (d *device) GetPropertyDriver() (string, error) {
	return d.getStringProperty(DevicePropertyDriver)
}

func (d *device) GetPropertyDriverVersion() (string, error) {
	return d.getStringProperty(DevicePropertyDriverVersion)
}

func (d *device) GetPropertyFirmwareVersion() (string, error) {
	return d.getStringProperty(DevicePropertyFirmwareVersion)
}

func (d *device) GetPropertyState() (NmDeviceState, error) {
	v, err := d.getUint32Property(DevicePropertyState)
	return NmDeviceState(v), err
}

func (d *device) GetPropertyIP4Config() (IP4Config, error) {
	path, err := d.getObjectProperty(DevicePropertyIp4Config)
	if err != nil || path == "/" {
		return nil, err
	}

	return NewIP4Config(path)
}

func (d *device) GetPropertyDHCP4Config() (DHCP4Config, error) {
	path, err := d.getObjectProperty(DevicePropertyDhcp4Config)
	if err != nil || path == "/" {
		return nil, err
	}

	return NewDHCP4Config(path)
}

func (d *device) GetPropertyIP6Config() (IP6Config, error) {
	path, err := d.getObjectProperty(DevicePropertyIp6Config)
	if err != nil || path == "/" {
		return nil, err
	}

	return NewIP6Config(path)
}

func (d *device) GetPropertyDHCP6Config() (DHCP6Config, error) {
	path, err := d.getObjectProperty(DevicePropertyDhcp6Config)
	if err != nil || path == "/" {
		return nil, err
	}

	return NewDHCP6Config(path)
}

func (d *device) GetPropertyManaged() (bool, error) {
	return d.getBoolProperty(DevicePropertyManaged)
}

func (d *device) GetPropertyAutoConnect() (bool, error) {
	return d.getBoolProperty(DevicePropertyAutoconnect)
}

func (d *device) GetPropertyFirmwareMissing() (bool, error) {
	return d.getBoolProperty(DevicePropertyFirmwareMissing)
}

func (d *device) GetPropertyNmPluginMissing() (bool, error) {
	return d.getBoolProperty(DevicePropertyNmPluginMissing)
}

func (d *device) GetPropertyDeviceType() (NmDeviceType, error) {
	v, err := d.getUint32Property(DevicePropertyDeviceType)
	return NmDeviceType(v), err
}

func (d *device) GetPropertyAvailableConnections() ([]Connection, error) {
	connPaths, err := d.getSliceObjectProperty(DevicePropertyAvailableConnections)
	if err != nil {
		return nil, err
	}

	conns := make([]Connection, len(connPaths))
	for i, path := range connPaths {
		conns[i], err = NewConnection(path)
		if err != nil {
			return conns, err
		}
	}

	return conns, nil
}

func (d *device) GetPropertyPhysicalPortId() (string, error) {
	return d.getStringProperty(DevicePropertyPhysicalPortId)
}

func (d *device) GetPropertyMtu() (uint32, error) {
	return d.getUint32Property(DevicePropertyMtu)
}

func (d *device) GetPropertyReal() (bool, error) {
	return d.getBoolProperty(DevicePropertyReal)
}

func (d *device) marshalMap() map[string]interface{} {
	m := make(map[string]interface{})

	m["Interface"], _ = d.GetPropertyInterface()
	m["IPInterface"], _ = d.GetPropertyIpInterface()
	m["IP4Config"], _ = d.GetPropertyIP4Config()
	m["DHCP4Config"], _ = d.GetPropertyDHCP4Config()
	m["AvailableConnections"], _ = d.GetPropertyAvailableConnections()

	state, _ := d.GetPropertyState()
	m["State"] = state.String()

	deviceType, _ := d.GetPropertyDeviceType()
	m["DeviceType"] = deviceType.String()

	return m
}

func (d *device) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.marshalMap())
}
