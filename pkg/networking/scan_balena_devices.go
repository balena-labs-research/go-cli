package networking

import (
	"context"
	"net"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/balena-community/go-cli/pkg/docker"
	"github.com/docker/docker/api/types"
)

func ScanBalenaDevices() ([]types.Info, error) {
	var deviceInfo []types.Info
	// Arp scan for available devices
	arpResult, err := ArpScan()

	if err != nil {
		return nil, err
	}

	// Scan to see if port 22222 is open as indicator of whether it is a balena device
	var wg sync.WaitGroup
	for _, ip := range *arpResult {
		wg.Add(1)

		// Run scans concurrently in goroutines
		go func(ip net.IP) {
			defer wg.Done()
			if err := ScanPort(ip.String(), 22222, time.Second*4); err == nil {
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

	return deviceInfo, nil
}
