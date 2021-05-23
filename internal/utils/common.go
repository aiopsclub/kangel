package utils

import (
	"github.com/docker/docker/client"
	"time"
)

func IsBlock(dockerClient *client.Client, containerID string, timeout time.Duration) bool {
	if Inspect(dockerClient, containerID, timeout) || Stats(dockerClient, containerID, timeout) {
		return true
	}
	return false
}
