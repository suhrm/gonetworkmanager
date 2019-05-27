package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus"
)

const (
	DHCP6ConfigInterface = NetworkManagerInterface + ".DHCP6Config"

	DHCP6ConfigPropertyOptions = DHCP6ConfigInterface + ".Options"
)

type DHCP6Options map[string]interface{}

type DHCP6Config interface {
	// GetOptions gets options map of configuration returned by the IPv4 DHCP server.
	GetOptions() DHCP6Options

	MarshalJSON() ([]byte, error)
}

func NewDHCP6Config(objectPath dbus.ObjectPath) (DHCP6Config, error) {
	var c dhcp6Config
	return &c, c.init(NetworkManagerInterface, objectPath)
}

type dhcp6Config struct {
	dbusBase
}

func (c *dhcp6Config) GetOptions() DHCP6Options {
	options := c.getMapStringVariantProperty(DHCP6ConfigPropertyOptions)
	rv := make(DHCP6Options)

	for k, v := range options {
		rv[k] = v.Value()
	}

	return rv
}

func (c *dhcp6Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"Options": c.GetOptions(),
	})
}
