package docker

import (
	"context"

	"github.com/docker/docker/api/types"
)

func ContainerList(ip string) ([]types.Container, error) {
	client, err := NewClient(ip)
	if err != nil {
		return nil, err
	}

	defer client.Close()

	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	return containers, err
}
