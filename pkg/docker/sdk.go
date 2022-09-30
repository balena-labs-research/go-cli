package docker

import (
	"log"

	"github.com/docker/docker/client"
)

// Create a new connection to the Docker daemon.
func NewClient(ip string, port string) *client.Client {

	tcpAddress := "tcp://" + ip + ":" + port
	cli, err := client.NewClientWithOpts(client.WithHost(tcpAddress), client.WithAPIVersionNegotiation())

	if err != nil {
		log.Fatal(err)
	}

	return cli
}
