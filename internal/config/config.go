package config

import (
	"encoding/json"
	"io"
	"os"

	"github.com/jerryagbesi/traffik/internal/server"
)

// LoadServers reads the server configuration from the given JSON file.
// helps the load balancer know which servers to proxy requests to
func LoadServers(filename string) ([]*server.Server, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var servers []*server.Server
	err = json.Unmarshal(content, &servers)
	if err != nil {
		return nil, err
	}

	return servers, nil
}
