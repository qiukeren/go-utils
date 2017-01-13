package utils

import (
	"github.com/fsouza/go-dockerclient"
	"os"
)

func IsDockerHost() bool {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		return false
	}
	return true
}

func IsDockerContainer() bool {
	if fileInfo, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}
