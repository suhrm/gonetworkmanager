package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus"
)

const (
	IP4ConfigInterface = NetworkManagerInterface + ".IP4Config"

	/* Properties */
	IP4ConfigPropertyAddresses      = IP4ConfigInterface + ".Addresses"      // readable   aau
	IP4ConfigPropertyAddressData    = IP4ConfigInterface + ".AddressData"    // readable   aa{sv}
	IP4ConfigPropertyGateway        = IP4ConfigInterface + ".Gateway"        // readable   s
	IP4ConfigPropertyRoutes         = IP4ConfigInterface + ".Routes"         // readable   aau
	IP4ConfigPropertyRouteData      = IP4ConfigInterface + ".RouteData"      // readable   aa{sv}
	IP4ConfigPropertyNameservers    = IP4ConfigInterface + ".Nameservers"    // readable   au
	IP4ConfigPropertyNameserverData = IP4ConfigInterface + ".NameserverData" // readable   aa{sv}
	IP4ConfigPropertyDomains        = IP4ConfigInterface + ".Domains"        // readable   as
	IP4ConfigPropertySearches       = IP4ConfigInterface + ".Searches"       // readable   as
	IP4ConfigPropertyDnsOptions     = IP4ConfigInterface + ".DnsOptions"     // readable   as
	IP4ConfigPropertyDnsPriority    = IP4ConfigInterface + ".DnsPriority"    // readable   i
	IP4ConfigPropertyWinsServers    = IP4ConfigInterface + ".WinsServers"    // readable   au
	IP4ConfigPropertyWinsServerData = IP4ConfigInterface + ".WinsServerData" // readable   as
)

// Deprecated: use IP4AddressData instead
type IP4Address struct {
	Address string
	Prefix  uint8
	Gateway string
}

type IP4AddressData struct {
	Address string
	Prefix  uint8
}

// Deprecated: use IP4RouteData instead
type IP4Route struct {
	Route   string
	Prefix  uint8
	NextHop string
	Metric  uint8
}

type IP4RouteData struct {
	Destination          string
	Prefix               uint8
	NextHop              string
	Metric               uint8
	AdditionalAttributes []string
}

type IP4NameserverData struct {
	Address string
}

type IP4Config interface {
	// Array of arrays of IPv4 address/prefix/gateway. All 3 elements of each array are in network byte order. Essentially: [(addr, prefix, gateway), (addr, prefix, gateway), ...]
	// Deprecated: use AddressData and Gateway
	GetAddresses() []IP4Address

	// Array of IP address data objects. All addresses will include "address" (an IP address string), and "prefix" (a uint). Some addresses may include additional attributes.
	GetAddressData() []IP4AddressData

	// The gateway in use.
	GetGateway() string

	// Arrays of IPv4 route/prefix/next-hop/metric. All 4 elements of each tuple are in network byte order. 'route' and 'next hop' are IPv4 addresses, while prefix and metric are simple unsigned integers. Essentially: [(route, prefix, next-hop, metric), (route, prefix, next-hop, metric), ...]
	// Deprecated: use RouteData
	GetRoutes() []IP4Route

	// Array of IP route data objects. All routes will include "dest" (an IP address string) and "prefix" (a uint). Some routes may include "next-hop" (an IP address string), "metric" (a uint), and additional attributes.
	GetRouteData() []IP4RouteData

	// The nameservers in use.
	// Deprecated: use NameserverData
	GetNameservers() []string

	// The nameservers in use. Currently only the value "address" is recognized (with an IP address string).
	GetNameserverData() []IP4NameserverData

	// A list of domains this address belongs to.
	GetDomains() []string

	// A list of dns searches.
	GetSearches() []string

	// A list of DNS options that modify the behavior of the DNS resolver. See resolv.conf(5) manual page for the list of supported options.
	GetDnsOptions() []string

	// The relative priority of DNS servers.
	GetDnsPriority() uint32

	// The Windows Internet Name Service servers associated with the connection.
	GetWinsServerData() []string

	MarshalJSON() ([]byte, error)
}

func NewIP4Config(objectPath dbus.ObjectPath) (IP4Config, error) {
	var c ip4Config
	return &c, c.init(NetworkManagerInterface, objectPath)
}

type ip4Config struct {
	dbusBase
}

// Deprecated: use GetAddressData
func (c *ip4Config) GetAddresses() []IP4Address {
	addresses := c.getSliceSliceUint32Property(IP4ConfigPropertyAddresses)
	ret := make([]IP4Address, len(addresses))

	for i, parts := range addresses {
		ret[i] = IP4Address{
			Address: ip4ToString(parts[0]),
			Prefix:  uint8(parts[1]),
			Gateway: ip4ToString(parts[2]),
		}
	}

	return ret
}

func (c *ip4Config) GetAddressData() []IP4AddressData {
	return []IP4AddressData{}
}

func (c *ip4Config) GetGateway() string {
	return c.getStringProperty(IP4ConfigPropertyGateway)
}

// Deprecated: use GetRouteData
func (c *ip4Config) GetRoutes() []IP4Route {
	routes := c.getSliceSliceUint32Property(IP4ConfigPropertyRoutes)
	ret := make([]IP4Route, len(routes))

	for i, parts := range routes {
		ret[i] = IP4Route{
			Route:   ip4ToString(parts[0]),
			Prefix:  uint8(parts[1]),
			NextHop: ip4ToString(parts[2]),
			Metric:  uint8(parts[3]),
		}
	}

	return ret
}

func (c *ip4Config) GetRouteData() []IP4RouteData {
	return []IP4RouteData{}
}

// Deprecated: use GetNameserverData
func (c *ip4Config) GetNameservers() []string {
	nameservers := c.getSliceUint32Property(IP4ConfigPropertyNameservers)
	ret := make([]string, len(nameservers))

	for i, ns := range nameservers {
		ret[i] = ip4ToString(ns)
	}

	return ret
}

func (c *ip4Config) GetNameserverData() []IP4NameserverData {
	return []IP4NameserverData{}
}

func (c *ip4Config) GetDomains() []string {
	return c.getSliceStringProperty(IP4ConfigPropertyDomains)
}

func (c *ip4Config) GetSearches() []string {
	return c.getSliceStringProperty(IP4ConfigPropertySearches)
}

func (c *ip4Config) GetDnsOptions() []string {
	return c.getSliceStringProperty(IP4ConfigPropertyDnsOptions)
}

func (c *ip4Config) GetDnsPriority() uint32 {
	return c.getUint32Property(IP4ConfigPropertyDnsPriority)
}

func (c *ip4Config) GetWinsServerData() []string {
	return c.getSliceStringProperty(IP4ConfigPropertyWinsServerData)
}

func (c *ip4Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"Addresses":   c.GetAddresses(),
		"Routes":      c.GetRoutes(),
		"Nameservers": c.GetNameservers(),
		"Domains":     c.GetDomains(),
	})
}
