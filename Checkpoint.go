package gonetworkmanager

import "github.com/godbus/dbus"

type Checkpoint interface {

	GetPath() dbus.ObjectPath
}

func NewCheckpoint(objectPath dbus.ObjectPath) (Checkpoint, error) {
	var c checkpoint
	return &c, c.init(NetworkManagerInterface, objectPath)
}

type checkpoint struct {
	dbusBase
}

func (c *checkpoint) GetPath() dbus.ObjectPath {
	return c.obj.Path()
}
