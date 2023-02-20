package networking

import (
	"context"
	"net"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/balena-labs-research/go-cli/pkg/docker"
	"github.com/docker/docker/api/types"
)

type DockerResponse struct {
	Info    types.Info `json:"info"`
	Address string     `json:"address"`
}

func CheckIpPorts(ips *[]net.IP, port int) []DockerResponse {
	var deviceInfo []DockerResponse

	// Scan to see if port is open as indicator of whether it is a balena device
	var wg sync.WaitGroup
	for _, ip := range *ips {
		wg.Add(1)

		// Run scans concurrently in goroutines
		go func(ip net.IP) {
			defer wg.Done()
			if err := DialPort(ip.String(), port, time.Second*4); err == nil {
				client, err := docker.NewClient(ip.String())
				if err != nil {
					log.Fatal(err)
				}

				defer client.Close()

				dClient, err := client.Info(context.Background())

				if err != nil {
					// Device is not accessible via the docker socket. Either not a
					// balena device, or is running in production more. Skipping.
					return
				}
				deviceInfo = append(deviceInfo, DockerResponse{Info: dClient, Address: ip.String()})

			}
		}(ip)
	}
	wg.Wait()

	return deviceInfo
}

func CheckHostnamePorts(hostnames []string, port int) []DockerResponse {
	var deviceInfo []DockerResponse

	// Scan to see if port is open as indicator of whether it is a balena device
	var wg sync.WaitGroup
	for _, hostname := range hostnames {
		wg.Add(1)

		// Run scans concurrently in goroutines
		go func(hostname string) {
			defer wg.Done()

			localHostname := hostname

			if err := DialPort(localHostname, port, time.Second*4); err == nil {
				client, err := docker.NewClient(localHostname)
				if err != nil {
					log.Fatal(err)
				}

				defer client.Close()

				dClient, err := client.Info(context.Background())

				if err != nil {
					// Device is not accessible via the docker socket. Either not a
					// balena device, or is running in production more. Skipping.
					return
				}

				deviceInfo = append(deviceInfo, DockerResponse{Info: dClient, Address: localHostname})

			}
		}(hostname)
	}
	wg.Wait()

	return deviceInfo
}
