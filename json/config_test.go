package json

import (
	"testing"
)

func TestConfig(t *testing.T) {
	confString := []byte(`
{
"MountPoint": "/mnt/test"
}
`)

	conf, err := Config(confString)

	if err != nil {
		t.Error(err)
	}

	if conf.MountPoint != "/mnt/test" {
		t.Error("Decoded bad config: MountPoint")
	}

}
