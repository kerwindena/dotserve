package json

import (
	"github.com/kerwindena/dotserve"

	"encoding/json"
)

// Config parses a byte string and returns the filled Config Object
func Config(b []byte) (dotserve.Config, error) {
	conf := dotserve.Config{}
	err := json.Unmarshal(b, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
