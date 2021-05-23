package utils

import (
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/shirou/gopsutil/v3/process"
	"strings"
)

func FindContainerPid(container types.Container) (*process.Process, error) {
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	for _, processIns := range processes {
		cmdline, err := processIns.Cmdline()
		if err != nil {
			return nil, err
		}
		if strings.Contains(cmdline, container.ID) {
			return processIns, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Pid of %s not found!", container.Names[0]))

}

func KillContainerPid(dockerClient *client.Client, container types.Container) error {
	process, findProcessErr := FindContainerPid(container)
	if findProcessErr != nil {
		return findProcessErr
	}
	killProcessErr := process.Kill()
	if killProcessErr != nil {
		return errors.New(fmt.Sprintf("kill %d for container %s error. error: %s", process.Pid, container.ID, killProcessErr))
	}
	return nil

}
