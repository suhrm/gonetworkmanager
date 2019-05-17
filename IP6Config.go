package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus"
)

const (
	IP6ConfigInterface = NetworkManagerInterface + ".IP6Config"

	/* Properties */
	IP6ConfigPropertyAddresses   = IP6ConfigInterface + ".Addresses"   // readable   a(ayuay)
	IP6ConfigPropertyAddressData = IP6ConfigInterface + ".AddressData" // readable   aa{sv}
	IP6ConfigPropertyGateway     = IP6ConfigInterface + ".Gateway"     // readable   s
	IP6ConfigPropertyRoutes      = IP6ConfigInterface + ".Routes"      // readable   a(ayuayu)
	IP6ConfigPropertyRouteData   = IP6ConfigInterface + ".RouteData"   // readable   aa{sv}
	IP6ConfigPropertyNameservers = IP6ConfigInterface + ".Nameservers" // readable   aay
	IP6ConfigPropertyDomains     = IP6ConfigInterface + ".Domains"     // readable   as
	IP6ConfigPropertySearches    = IP6ConfigInterface + ".Searches"    // readable   as
	IP6ConfigPropertyDnsOptions  = IP6ConfigInterface + ".DnsOptions"  // readable   as
	IP6ConfigPropertyDnsPriority = IP6ConfigInterface + ".DnsPriority" // readable   i
)

// Deprecated: use IP6AddressData instead
type IP6Address struct {
	Address string
	Prefix  uint8
	Gateway string
}

type IP6AddressData struct {
	Address string
	Prefix  uint8
}

// Deprecated: use IP6RouteData instead
type IP6Route struct {
	Route   string
	Prefix  uint8
	NextHop string
	Metric  uint8
}

type IP6RouteData struct {
	Destination          string
	Prefix               uint8
	NextHop              string
	Metric               uint8
	AdditionalAttributes []string
}

type IP6NameserverData struct {
	Address string
}

type IP6Config interface {

	// Array of IP address data objects. All addresses will include "address" (an IP address string), and "prefix" (a uint). Some addresses may include additional attributes.
	GetAddressData() []IP6AddressData

	// The gateway in use.
	GetGateway() string

	// Array of IP route data objects. All routes will include "dest" (an IP address string) and "prefix" (a uint). Some routes may include "next-hop" (an IP address string), "metric" (a uint), and additional attributes.
	GetRouteData() []IP6RouteData

	// GetNameservers gets the nameservers in use.
	GetNameservers() []IP6NameserverData

	// A list of domains this address belongs to.
	GetDomains() []string

	// A list of dns searches.
	GetSearches() []string

	// A list of DNS options that modify the behavior of the DNS resolver. See resolv.conf(5) manual page for the list of supported options.
	GetDnsOptions() []string

	// The relative priority of DNS servers.
	GetDnsPriority() uint32

	MarshalJSON() ([]byte, error)
}

func NewIP6Config(objectPath dbus.ObjectPath) (IP6Config, error) {
	var c ip6Config
	return &c, c.init(NetworkManagerInterface, objectPath)
}

type ip6Config struct {
	dbusBase
}

func (c *ip6Config) GetAddressData() []IP6AddressData {
	return []IP6AddressData{}
}

func (c *ip6Config) GetGateway() string {
	return c.getStringProperty(IP6ConfigPropertyGateway)
}

func (c *ip6Config) GetRouteData() []IP6RouteData {
	return []IP6RouteData{}
}

func (c *ip6Config) GetNameservers() []IP6NameserverData {
	return []IP6NameserverData{}
}

func (c *ip6Config) GetDomains() []string {
	return c.getSliceStringProperty(IP6ConfigPropertyDomains)
}

func (c *ip6Config) GetSearches() []string {
	return c.getSliceStringProperty(IP6ConfigPropertySearches)
}

func (c *ip6Config) GetDnsOptions() []string {
	return c.getSliceStringProperty(IP6ConfigPropertyDnsOptions)
}

func (c *ip6Config) GetDnsPriority() uint32 {
	return c.getUint32Property(IP6ConfigPropertyDnsPriority)
}

func (c *ip6Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"Addresses":   c.GetAddressData(),
		"Routes":      c.GetRouteData(),
		"Nameservers": c.GetNameservers(),
		"Domains":     c.GetDomains(),
	})
}
