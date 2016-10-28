package json

import (
	"encoding/json"
	"sync"
)

// confData actually stores the data and is used by unmarshaling
type confData struct {
	MountPoint string
}

// Config can read json input and stores the data internally
type Config struct {
	rwLock sync.RWMutex
	data   confData
}

// MountPoint gets the mountpoint from the internal structure
func (c *Config) MountPoint() string {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()

	return c.data.MountPoint
}

// Update reads a bytestring as json and updates its internal data storage
func (c *Config) Update(b []byte) error {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()

	conf := &confData{}
	err := json.Unmarshal(b, conf)
	if err != nil {
		return err
	}

	c.data.MountPoint = conf.MountPoint

	return nil
}
