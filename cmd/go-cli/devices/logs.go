package devices

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/balena-labs-research/go-cli/pkg/docker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/stdcopy"
)

func StreamAllLogs(args []string) {
	var address string

	if len(args) > 0 {
		address = args[0]
	} else {
		index, deviceInfo := selectLocalDevice()

		if len(deviceInfo) == 0 {
			fmt.Println("No devices found")
			return
		}

		address = getLocalDeviceAddress(index, deviceInfo)
	}

	client, err := docker.NewClient(address)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {
		log.Fatal(err)
	}

	// Check to see if only the Supervisor container is running (i.e. no user containers)
	if len(containers) == 1 {
		fmt.Println("No containers running")
		return
	}

	var wg sync.WaitGroup

	for _, container := range containers {
		if container.Names[0] != "/balena_supervisor" {
			wg.Add(1)
			go func(container types.Container) {
				defer wg.Done()
				out, err := client.ContainerLogs(context.Background(), container.ID, types.ContainerLogsOptions{
					ShowStdout: true,
					ShowStderr: true,
					Follow:     true,
					Tail:       "0"})

				if err != nil {
					log.Error(err)
					return
				}

				fmt.Println("[balena-CLI Logs] Streaming logs for " + container.Names[0])

				// Depending on TTY setting on container it may contain headers and needs stripping
				_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, out)

				// When not TTY, io is required instead
				if err != nil {
					_, err = io.Copy(os.Stdout, out)
					if err != nil && err != io.EOF {
						log.Fatal(err)
					}
				}

				fmt.Println("[balena-CLI Logs] Disconnected from " + container.Names[0])
			}(container)
		}
	}

	wg.Wait()

	fmt.Println("[balena-CLI Logs] Disconnected from all containers")
}
