package json

import (
	"testing"
)

var (
	configString1 = []byte(`
{
"MountPoint": "/mnt/test"
}
`)
)

func TestConfig(t *testing.T) {
	conf := &Config{}
	err := conf.Update(configString1)

	if err != nil {
		t.Error(err)
	}
}

func TestMountPoint(t *testing.T) {
	conf := &Config{}
	err := conf.Update(configString1)

	if err != nil {
		t.Error(err)
	}

	if conf.MountPoint() != "/mnt/test" {
		t.Error("Decoded bad config: MountPoint")
	}
}
