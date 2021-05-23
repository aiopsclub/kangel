package docker

import (
	"github.com/docker/docker/client"
	"kangel/internal/utils"
	"time"
)

func BlockDetect() error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	containers, containerPsErr := utils.ContainerPs(cli, time.Second*3)

	if containerPsErr != nil {
		return containerPsErr
	}

	for _, container := range containers {
		if utils.IsBlock(cli, container.ID, time.Second*10) {
			err := utils.KillContainerPid(cli, container)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
