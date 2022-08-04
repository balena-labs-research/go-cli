package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Return info from client
func Info(cli *client.Client) (types.Info, error) {
	ctx := context.Background()

	info, err := cli.Info(ctx)

	if err != nil {
		return info, err
	}

	return info, err
}

// Create a new connection to the Docker daemon.
func NewClient(ip string, port string) (*client.Client, error) {

	tcpAddress := "tcp://" + ip + ":" + port
	cli, err := client.NewClientWithOpts(client.WithHost(tcpAddress), client.WithAPIVersionNegotiation())

	if err != nil {
		return cli, err
	}
	return cli, err
}
