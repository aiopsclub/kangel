package utils

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ContainerInspect(dockerClient *client.Client, container types.Container, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_, err := dockerClient.ContainerInspect(ctx, container.ID)
	if errors.Is(err, context.DeadlineExceeded) {
		log.Printf("%s is blocking...", container.ID)
		return true
	}
	return false
}

func ContainerStats(dockerClient *client.Client, container types.Container, timeout time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_, err := dockerClient.ContainerStats(ctx, container.ID, false)
	if errors.Is(err, context.DeadlineExceeded) {
		log.Printf("%s is blocking...", container.ID)
		return true
	}
	return false
}

func ContainerPs(dockerClient *client.Client, timeout time.Duration) ([]types.Container, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	containerList, err := dockerClient.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return containerList, err
	}
	return containerList, nil
}
