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

func CheckIpPorts(ips *[]net.IP, port int) []types.Info {
	var deviceInfo []types.Info

	// Scan to see if port is open as indicator of whether it is a balena device
	var wg sync.WaitGroup
	for _, ip := range *ips {
		wg.Add(1)

		// Run scans concurrently in goroutines
		go func(ip net.IP) {
			defer wg.Done()
			if err := ScanPort(ip.String(), port, time.Second*4); err == nil {
				client := docker.NewClient(ip.String(), "2375")

				if err != nil {
					log.Error(err)
				}

				dClient, err := client.Info(context.Background())

				if err != nil {
					// Device is not accessible via the docker socket. Either not a
					// balena device, or is running in production more. Skipping.
					return
				}

				deviceInfo = append(deviceInfo, dClient)
			}
		}(ip)
	}
	wg.Wait()

	return deviceInfo
}

func CheckHostnamePorts(hostnames []string, port int) []types.Info {
	var deviceInfo []types.Info

	// Scan to see if port is open as indicator of whether it is a balena device
	var wg sync.WaitGroup
	for _, hostname := range hostnames {
		wg.Add(1)

		// Run scans concurrently in goroutines
		go func(hostname string) {
			defer wg.Done()

			localHostname := hostname

			if err := ScanPort(localHostname, port, time.Second*4); err == nil {
				client := docker.NewClient(localHostname, "2375")

				if err != nil {
					log.Error(err)
				}

				dClient, err := client.Info(context.Background())

				if err != nil {
					// Device is not accessible via the docker socket. Either not a
					// balena device, or is running in production more. Skipping.
					return
				}

				deviceInfo = append(deviceInfo, dClient)
			}
		}(hostname)
	}
	wg.Wait()

	return deviceInfo
}
