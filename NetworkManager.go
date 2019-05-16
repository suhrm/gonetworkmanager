package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus"
)

const (
	NetworkManagerInterface  = "org.freedesktop.NetworkManager"
	NetworkManagerObjectPath = "/org/freedesktop/NetworkManager"

	/* Methods */
	NetworkManagerReload                          = NetworkManagerInterface + ".Reload"
	NetworkManagerGetDevices                      = NetworkManagerInterface + ".GetDevices"
	NetworkManagerGetAllDevices                   = NetworkManagerInterface + ".GetAllDevices"
	NetworkManagerGetDeviceByIpIface              = NetworkManagerInterface + ".GetDeviceByIpIface"
	NetworkManagerActivateConnection              = NetworkManagerInterface + ".ActivateConnection"
	NetworkManagerAddAndActivateConnection        = NetworkManagerInterface + ".AddAndActivateConnection"
	NetworkManagerAddAndActivateConnection2       = NetworkManagerInterface + ".AddAndActivateConnection2"
	NetworkManagerDeactivateConnection            = NetworkManagerInterface + ".DeactivateConnection"
	NetworkManagerSleep                           = NetworkManagerInterface + ".Sleep"
	NetworkManagerEnable                          = NetworkManagerInterface + ".Enable"
	NetworkManagerGetPermissions                  = NetworkManagerInterface + ".GetPermissions"
	NetworkManagerSetLogging                      = NetworkManagerInterface + ".SetLogging"
	NetworkManagerGetLogging                      = NetworkManagerInterface + ".GetLogging"
	NetworkManagerCheckConnectivity               = NetworkManagerInterface + ".CheckConnectivity"
	NetworkManagerstate                           = NetworkManagerInterface + ".state"
	NetworkManagerCheckpointCreate                = NetworkManagerInterface + ".CheckpointCreate"
	NetworkManagerCheckpointDestroy               = NetworkManagerInterface + ".CheckpointDestroy"
	NetworkManagerCheckpointRollback              = NetworkManagerInterface + ".CheckpointRollback"
	NetworkManagerCheckpointAdjustRollbackTimeout = NetworkManagerInterface + ".CheckpointAdjustRollbackTimeout"

	/* Property */
	NetworkManagerPropertyDevices                    = NetworkManagerInterface + ".Devices"                    // readable   ao
	NetworkManagerPropertyAllDevices                 = NetworkManagerInterface + ".AllDevices"                 // readable   ao
	NetworkManagerPropertyCheckpoints                = NetworkManagerInterface + ".Checkpoints"                // readable   ao
	NetworkManagerPropertyNetworkingEnabled          = NetworkManagerInterface + ".NetworkingEnabled"          // readable   b
	NetworkManagerPropertyWirelessEnabled            = NetworkManagerInterface + ".WirelessEnabled"            // readwrite  b
	NetworkManagerPropertyWirelessHardwareEnabled    = NetworkManagerInterface + ".WirelessHardwareEnabled"    // readable   b
	NetworkManagerPropertyWwanEnabled                = NetworkManagerInterface + ".WwanEnabled"                // readwrite  b
	NetworkManagerPropertyWwanHardwareEnabled        = NetworkManagerInterface + ".WwanHardwareEnabled"        // readable   b
	NetworkManagerPropertyWimaxEnabled               = NetworkManagerInterface + ".WimaxEnabled"               // readwrite  b
	NetworkManagerPropertyWimaxHardwareEnabled       = NetworkManagerInterface + ".WimaxHardwareEnabled"       // readable   b
	NetworkManagerPropertyActiveConnections          = NetworkManagerInterface + ".ActiveConnections"          // readable   ao
	NetworkManagerPropertyPrimaryConnection          = NetworkManagerInterface + ".PrimaryConnection"          // readable   o
	NetworkManagerPropertyPrimaryConnectionType      = NetworkManagerInterface + ".PrimaryConnectionType"      // readable   s
	NetworkManagerPropertyMetered                    = NetworkManagerInterface + ".Metered"                    // readable   u
	NetworkManagerPropertyActivatingConnection       = NetworkManagerInterface + ".ActivatingConnection"       // readable   o
	NetworkManagerPropertyStartup                    = NetworkManagerInterface + ".Startup"                    // readable   b
	NetworkManagerPropertyVersion                    = NetworkManagerInterface + ".Version"                    // readable   s
	NetworkManagerPropertyCapabilities               = NetworkManagerInterface + ".Capabilities"               // readable   au
	NetworkManagerPropertyState                      = NetworkManagerInterface + ".State"                      // readable   u
	NetworkManagerPropertyConnectivity               = NetworkManagerInterface + ".Connectivity"               // readable   u
	NetworkManagerPropertyConnectivityCheckAvailable = NetworkManagerInterface + ".ConnectivityCheckAvailable" // readable   b
	NetworkManagerPropertyConnectivityCheckEnabled   = NetworkManagerInterface + ".ConnectivityCheckEnabled"   // readwrite  b
	NetworkManagerPropertyGlobalDnsConfiguration     = NetworkManagerInterface + ".GlobalDnsConfiguration"     // readwrite  a{sv}
)

type NetworkManager interface {
	// Get the list of realized network devices.
	GetDevices() []Device

	// Get the list of all network devices.
	GetAllDevices() []Device

	// GetState returns the overall networking state as determined by the
	// NetworkManager daemon, based on the state of network devices under it's
	// management.
	GetState() NmState

	// GetActiveConnections returns the active connection of network devices.
	GetActiveConnections() []ActiveConnection

	// ActivateWirelessConnection requests activating access point to network device
	ActivateWirelessConnection(connection Connection, device Device, accessPoint AccessPoint) ActiveConnection

	// AddAndActivateWirelessConnection adds a new connection profile to the network device it has been
	// passed. It then activates the connection to the passed access point. The first paramter contains
	// additional information for the connection (most propably the credentials).
	// Example contents for connection are:
	// connection := make(map[string]map[string]interface{})
	// connection["802-11-wireless"] = make(map[string]interface{})
	// connection["802-11-wireless"]["security"] = "802-11-wireless-security"
	// connection["802-11-wireless-security"] = make(map[string]interface{})
	// connection["802-11-wireless-security"]["key-mgmt"] = "wpa-psk"
	// connection["802-11-wireless-security"]["psk"] = password
	AddAndActivateWirelessConnection(connection map[string]map[string]interface{}, device Device, accessPoint AccessPoint) (ac ActiveConnection, err error)

	Subscribe() <-chan *dbus.Signal
	Unsubscribe()

	MarshalJSON() ([]byte, error)
}

func NewNetworkManager() (NetworkManager, error) {
	var nm networkManager
	return &nm, nm.init(NetworkManagerInterface, NetworkManagerObjectPath)
}

type networkManager struct {
	dbusBase

	sigChan chan *dbus.Signal
}

func (n *networkManager) GetDevices() []Device {
	var devicePaths []dbus.ObjectPath

	n.callAndPanic(&devicePaths, NetworkManagerGetDevices)
	devices := make([]Device, len(devicePaths))

	var err error
	for i, path := range devicePaths {
		devices[i], err = DeviceFactory(path)
		if err != nil {
			panic(err)
		}
	}

	return devices
}

func (n *networkManager) GetAllDevices() []Device {
	var devicePaths []dbus.ObjectPath

	n.callAndPanic(&devicePaths, NetworkManagerGetAllDevices)
	devices := make([]Device, len(devicePaths))

	var err error
	for i, path := range devicePaths {
		devices[i], err = DeviceFactory(path)
		if err != nil {
			panic(err)
		}
	}

	return devices
}

func (n *networkManager) GetState() NmState {
	return NmState(n.getUint32Property(NetworkManagerPropertyState))
}

func (n *networkManager) GetActiveConnections() []ActiveConnection {
	acPaths := n.getSliceObjectProperty(NetworkManagerPropertyActiveConnections)
	ac := make([]ActiveConnection, len(acPaths))

	var err error
	for i, path := range acPaths {
		ac[i], err = NewActiveConnection(path)
		if err != nil {
			panic(err)
		}
	}

	return ac
}

func (n *networkManager) ActivateWirelessConnection(c Connection, d Device, ap AccessPoint) ActiveConnection {
	var opath dbus.ObjectPath
	n.callAndPanic(&opath, NetworkManagerActivateConnection, c.GetPath(), d.GetPath(), ap.GetPath())
	return nil
}

func (n *networkManager) AddAndActivateWirelessConnection(connection map[string]map[string]interface{}, d Device, ap AccessPoint) (ac ActiveConnection, err error) {
	var opath1 dbus.ObjectPath
	var opath2 dbus.ObjectPath

	err = n.callWithReturn2(&opath1, &opath2, NetworkManagerAddAndActivateConnection, connection, d.GetPath(), ap.GetPath())
	if err != nil {
		return
	}

	ac, err = NewActiveConnection(opath2)
	if err != nil {
		return
	}
	return
}

func (n *networkManager) Subscribe() <-chan *dbus.Signal {
	if n.sigChan != nil {
		return n.sigChan
	}

	n.subscribeNamespace(NetworkManagerObjectPath)
	n.sigChan = make(chan *dbus.Signal, 10)
	n.conn.Signal(n.sigChan)

	return n.sigChan
}

func (n *networkManager) Unsubscribe() {
	n.conn.RemoveSignal(n.sigChan)
	n.sigChan = nil
}

func (n *networkManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"NetworkState": n.GetState().String(),
		"Devices":      n.GetDevices(),
	})
}
