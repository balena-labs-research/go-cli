package devices

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"golang.org/x/exp/slices"

	"github.com/balena-labs-research/go-cli/pkg/docker"
	"github.com/balena-labs-research/go-cli/pkg/networking"
	"github.com/balena-labs-research/go-cli/pkg/spinner"
)

func Scan(scanType string, args []string, listContainers bool) {
	var err error
	var deviceInfo []networking.DockerResponse
	switch scanType {
	case "arp":
		deviceInfo = arpScan()
	case "lookup":
		var ipRange string

		if len(args) > 0 {
			ipRange = args[0]
		}
		deviceInfo, err = lookupScan(ipRange)

		if err != nil {
			log.Print(err)
		}
	}

	// If no devices were found, return
	if len(deviceInfo) == 0 {
		fmt.Println("No devices found")
		return
	}

	// Print the info for each device
	for deviceNumber, item := range deviceInfo {
		fmt.Println("\033[32m- Device", deviceNumber+1, "-", "\033[0m")
		fmt.Printf("Address: %s \n", item.Address)
		fmt.Printf("Hostname: %s \n", item.Info.Name)
		if listContainers {
			fmt.Println("Containers: ")

			listOfContainers, _ := docker.ContainerList("100.121.162.79")

			for _, container := range listOfContainers {
				if slices.Contains(container.Names, "/balena_supervisor") {
					// Skipping the Supervisor container
					continue
				}

				fmt.Println("\t  Name:", container.Names[0][1:])
				fmt.Println("\t\tState:", container.State)
				fmt.Println("\t\tStatus:", container.Status)
			}
		} else {
			// Number of containers minus the Supervisor container
			fmt.Printf("Number of containers: %v \n", item.Info.Containers-1)
			fmt.Printf("Containers running: %v \n", item.Info.ContainersRunning-1)
		}
		fmt.Printf("Kernel Version: %s \n", item.Info.KernelVersion)
		fmt.Printf("Operating System: %s \n", item.Info.OperatingSystem)
		fmt.Printf("Architecture: %s \n", item.Info.Architecture)
		fmt.Printf("Balena Engine Version: %s \n", item.Info.ServerVersion)
		fmt.Printf(" \n")
	}
}

func arpScan() []networking.DockerResponse {
	s := spinner.StartNew("Scanning for local balenaOS devices...")

	// Arp scan for available devices
	arpResults, err := networking.ArpScan()

	if err != nil {
		log.Fatal(err)
	}

	deviceInfo := networking.CheckIpPorts(arpResults, 22222)

	// Stop before error to avoid overlap of messages
	s.Stop()

	return deviceInfo
}

func lookupScan(ipRange string) ([]networking.DockerResponse, error) {
	s := spinner.StartNew("Scanning for local balenaOS devices...")
	lookupResults, err := networking.LookupAddresses(ipRange)

	if err != nil {
		log.Fatal(err)
	}

	deviceInfo := networking.CheckHostnamePorts(lookupResults, 22222)

	// Stop before error to avoid overlap of messages
	s.Stop()

	return deviceInfo, nil
}
