package gonetworkmanager

import (
	"encoding/json"

	"github.com/godbus/dbus"
)

const (
	ConnectionInterface = SettingsInterface + ".Connection"

	/* Methods */
	ConnectionUpdate        = ConnectionInterface + ".Update"
	ConnectionUpdateUnsaved = ConnectionInterface + ".UpdateUnsaved"
	ConnectionDelete        = ConnectionInterface + ".Delete"
	ConnectionGetSettings   = ConnectionInterface + ".GetSettings"
	ConnectionGetSecrets    = ConnectionInterface + ".GetSecrets"
	ConnectionClearSecrets  = ConnectionInterface + ".ClearSecrets"
	ConnectionSave          = ConnectionInterface + ".Save"
	ConnectionUpdate2       = ConnectionInterface + ".Update2"

	/* Properties */
	ConnectionPropertyUnsaved  = ConnectionInterface + ".Unsaved"  // readable   b
	ConnectionPropertyFlags    = ConnectionInterface + ".Flags"    // readable   u
	ConnectionPropertyFilename = ConnectionInterface + ".Filename" // readable   s
)

//type ConnectionSettings map[string]map[string]interface{}
type ConnectionSettings map[string]map[string]interface{}

type Connection interface {
	GetPath() dbus.ObjectPath

	// Update the connection with new settings and properties (replacing all previous settings and properties) and save the connection to disk. Secrets may be part of the update request, and will be either stored in persistent storage or sent to a Secret Agent for storage, depending on the flags associated with each secret.
	Update(settings ConnectionSettings) error

	// Update the connection with new settings and properties (replacing all previous settings and properties) but do not immediately save the connection to disk. Secrets may be part of the update request and may sent to a Secret Agent for storage, depending on the flags associated with each secret. Use the 'Save' method to save these changes to disk. Note that unsaved changes will be lost if the connection is reloaded from disk (either automatically on file change or due to an explicit ReloadConnections call).
	UpdateUnsaved(settings ConnectionSettings) error

	// Delete the connection.
	Delete() error

	// GetSettings gets the settings maps describing this network configuration.
	// This will never include any secrets required for connection to the
	// network, as those are often protected. Secrets must be requested
	// separately using the GetSecrets() callWithReturnAndPanic.
	GetSettings() ConnectionSettings

	// Clear the secrets belonging to this network connection profile.
	ClearSecrets() error

	// Saves a "dirty" connection (that had previously been updated with UpdateUnsaved) to persistent storage.
	Save() error

	// If set, indicates that the in-memory state of the connection does not match the on-disk state. This flag will be set when UpdateUnsaved() is called or when any connection details change, and cleared when the connection is saved to disk via Save() or from internal operations.
	GetUnsaved() bool

	// Additional flags of the connection profile.
	GetFlags() uint32

	// File that stores the connection in case the connection is file-backed.
	GetFilename() string

	MarshalJSON() ([]byte, error)
}

func NewConnection(objectPath dbus.ObjectPath) (Connection, error) {
	var c connection
	return &c, c.init(NetworkManagerInterface, objectPath)
}

type connection struct {
	dbusBase
}

func (c *connection) GetPath() dbus.ObjectPath {
	return c.obj.Path()
}

func (c *connection) Update(settings ConnectionSettings) error {
	return c.call(ConnectionUpdate, settings)
}

func (c *connection) UpdateUnsaved(settings ConnectionSettings) error {
	return c.call(ConnectionUpdateUnsaved, settings)
}

func (c *connection) Delete() error {
	return c.call(ConnectionDelete)
}

func (c *connection) GetSettings() ConnectionSettings {
	var settings map[string]map[string]dbus.Variant
	c.callWithReturnAndPanic(&settings, ConnectionGetSettings)

	rv := make(ConnectionSettings)

	for k1, v1 := range settings {
		rv[k1] = make(map[string]interface{})

		for k2, v2 := range v1 {
			rv[k1][k2] = v2.Value()
		}
	}

	return rv
}

func (c *connection) ClearSecrets() error {
	return c.call(ConnectionClearSecrets)
}

func (c *connection) Save() error {
	return c.call(ConnectionSave)
}

func (c *connection) GetUnsaved() bool {
	return c.getBoolProperty(ConnectionPropertyUnsaved)
}

func (c *connection) GetFlags() uint32 {
	return c.getUint32Property(ConnectionPropertyFlags)
}

func (c *connection) GetFilename() string {
	return c.getStringProperty(ConnectionPropertyFilename)
}

func (c *connection) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.GetSettings())
}
