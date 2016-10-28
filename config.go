package dotserve

// Config parses configuration data and stores all configuration
type Config interface {
	// MountPoint gets the mountpoint from the internal storage
	MountPoint() string
	// Update parses an bytestring into the internal storage
	Update([]byte) error
}
