package gonetworkmanager

import (
	"github.com/godbus/dbus"
)

const (
	SettingsInterface  = NetworkManagerInterface + ".Settings"
	SettingsObjectPath = NetworkManagerObjectPath + "/Settings"

	SettingsListConnections = SettingsInterface + ".ListConnections"
	SettingsAddConnection   = SettingsInterface + ".AddConnection"
	SettingsSaveHostname    = SettingsInterface + ".SaveHostname"

	SettingsHostnameProperty = SettingsInterface + ".Hostname"
)

type Settings interface {
	// ListConnections gets list the saved network connections known to NetworkManager
	ListConnections() []Connection

	// AddConnection callAndPanic new connection and save it to disk.
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

	s.callAndPanic(&connectionPaths, SettingsListConnections)
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
	s.callAndPanic(&path, SettingsAddConnection, settings)
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
	hostname := s.getStringProperty(SettingsHostnameProperty)

	return hostname
}
