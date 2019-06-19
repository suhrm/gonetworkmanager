package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus"
)

const (
	DeviceStatisticsInterface = DeviceInterface + ".Statistics"

	DeviceStatisticsPropertyRefreshRateMs = DeviceStatisticsInterface + ".RefreshRateMs" // readwrite  u
	DeviceStatisticsPropertyTxBytes       = DeviceStatisticsInterface + ".TxBytes"       // readable   t
	DeviceStatisticsPropertyRxBytes       = DeviceStatisticsInterface + ".RxBytes"       // readable   t
)

type DeviceStatistics interface {
	GetPath() dbus.ObjectPath

	// Refresh rate of the rest of properties of this interface. The properties are guaranteed to be refreshed each RefreshRateMs milliseconds in case the underlying counter has changed too. If zero, there is no guaranteed refresh rate of the properties.
	GetPropertyRefreshRateMs() uint32

	// Number of transmitted bytes
	GetPropertyTxBytes() uint64

	// Number of received bytes
	GetPropertyRxBytes() uint64
}

func NewDeviceStatistics(objectPath dbus.ObjectPath) (DeviceStatistics, error) {
	var d deviceStatistics
	return &d, d.init(NetworkManagerInterface, objectPath)
}

type deviceStatistics struct {
	dbusBase
}

func (d *deviceStatistics) GetPath() dbus.ObjectPath {
	return d.obj.Path()
}

func (d *deviceStatistics) GetPropertyRefreshRateMs() uint32 {
	return d.getUint32Property(DeviceStatisticsPropertyRefreshRateMs)
}

func (d *deviceStatistics) GetPropertyTxBytes() uint64 {
	return d.getUint64Property(DeviceStatisticsPropertyTxBytes)
}

func (d *deviceStatistics) GetPropertyRxBytes() uint64 {
	return d.getUint64Property(DeviceStatisticsPropertyRxBytes)
}

func (d *deviceStatistics) marshalMap() map[string]interface{} {
	return map[string]interface{}{
		"RefreshRateMs": d.GetPropertyRefreshRateMs(),
		"TxBytes": d.GetPropertyTxBytes(),
		"RxBytes": d.GetPropertyRxBytes(),
	}
}

func (d *deviceStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.marshalMap())
}
