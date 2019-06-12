package gonetworkmanager

import (
	"encoding/json"
	"github.com/godbus/dbus"
)

const (
	CheckpointInterface = NetworkManagerInterface + ".Checkpoint"

	/* Properties */
	CheckpointPropertyDevices         = CheckpointInterface + ".Devices"         // readable   ao
	CheckpointPropertyCreated         = CheckpointInterface + ".Created"         // readable   x
	CheckpointPropertyRollbackTimeout = CheckpointInterface + ".RollbackTimeout" // readable   u
)

type Checkpoint interface {
	GetPath() dbus.ObjectPath

	// Array of object paths for devices which are part of this checkpoint.
	GetPropertyDevices() []Device

	// The timestamp (in CLOCK_BOOTTIME milliseconds) of checkpoint creation.
	GetPropertyCreated() int64

	// Timeout in seconds for automatic rollback, or zero.
	GetPropertyRollbackTimeout() uint32

	MarshalJSON() ([]byte, error)
}

func NewCheckpoint(objectPath dbus.ObjectPath) (Checkpoint, error) {
	var c checkpoint
	return &c, c.init(NetworkManagerInterface, objectPath)
}

type checkpoint struct {
	dbusBase
}

func (c *checkpoint) GetPropertyDevices() []Device {
	devicesPaths := c.getSliceObjectProperty(CheckpointPropertyDevices)
	devices := make([]Device, len(devicesPaths))

	var err error
	for i, path := range devicesPaths {
		devices[i], err = NewDevice(path)
		if err != nil {
			panic(err)
		}
	}

	return devices
}

func (c *checkpoint) GetPropertyCreated() int64 {
	return c.getInt64Property(CheckpointPropertyCreated)
}

func (c *checkpoint) GetPropertyRollbackTimeout() uint32 {
	return c.getUint32Property(CheckpointPropertyRollbackTimeout)
}

func (c *checkpoint) GetPath() dbus.ObjectPath {
	return c.obj.Path()
}

func (c *checkpoint) marshalMap() map[string]interface{} {
	return map[string]interface{}{
		"Devices":         c.GetPropertyDevices(),
		"Created":         c.GetPropertyCreated(),
		"RollbackTimeout": c.GetPropertyRollbackTimeout(),
	}
}

func (c *checkpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.marshalMap())
}
