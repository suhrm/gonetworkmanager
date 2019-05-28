package gonetworkmanager

import (
	"github.com/godbus/dbus"
)

const (
	SettingsInterface  = NetworkManagerInterface + ".Settings"
	SettingsObjectPath = NetworkManagerObjectPath + "/Settings"

	/* Methods */
	SettingsListConnections      = SettingsInterface + ".ListConnections"
	SettingsGetConnectionByUuid  = SettingsInterface + ".GetConnectionByUuid"
	SettingsAddConnection        = SettingsInterface + ".AddConnection"
	SettingsAddConnectionUnsaved = SettingsInterface + ".AddConnectionUnsaved"
	SettingsLoadConnections      = SettingsInterface + ".LoadConnections"
	SettingsReloadConnections    = SettingsInterface + ".ReloadConnections"
	SettingsSaveHostname         = SettingsInterface + ".SaveHostname"

	/* Properties */
	SettingsPropertyConnections = SettingsInterface + ".Connections" // readable   ao
	SettingsPropertyHostname    = SettingsInterface + ".Hostname"    // readable   s
	SettingsPropertyCanModify   = SettingsInterface + ".CanModify"   // readable   b
)

type Settings interface {
	// ListConnections gets list the saved network connections known to NetworkManager
	ListConnections() []Connection

	// AddConnection callWithReturnAndPanic new connection and save it to disk.
	AddConnection(settings ConnectionSettings) Connection

	// Save the hostname to persistent configuration.
	SaveHostname(hostname string)

	// The machine hostname stored in persistent configuration.
	Hostname() string
}

func NewSettings() (Settings, error) {
	var s settings
	return &s, s.init(NetworkManagerInterface, SettingsObjectPath)
}

type settings struct {
	dbusBase
}

func (s *settings) ListConnections() []Connection {
	var connectionPaths []dbus.ObjectPath

	s.callWithReturnAndPanic(&connectionPaths, SettingsListConnections)
	connections := make([]Connection, len(connectionPaths))

	var err error
	for i, path := range connectionPaths {
		connections[i], err = NewConnection(path)
		if err != nil {
			panic(err)
		}
	}

	return connections
}

func (s *settings) AddConnection(settings ConnectionSettings) Connection {
	var path dbus.ObjectPath
	s.callWithReturnAndPanic(&path, SettingsAddConnection, settings)
	con, err := NewConnection(path)
	if err != nil {
		panic(err)
	}
	return con
}

func (s *settings) SaveHostname(hostname string) {
	err := s.call(SettingsSaveHostname, hostname)
	if err != nil {
		panic(err)
	}
}

func (s *settings) Hostname() string {
	hostname := s.getStringProperty(SettingsPropertyHostname)

	return hostname
}
