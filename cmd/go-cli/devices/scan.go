package devices

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/balena-labs-research/go-cli/pkg/networking"
	"github.com/balena-labs-research/go-cli/pkg/spinner"
)

func Scan(scanType string, args []string) {
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
		fmt.Printf(" - Device %v - \n", deviceNumber+1)
		fmt.Printf("Address: %s \n", item.Address)
		fmt.Printf("Hostname: %s \n", item.Info.Name)
		fmt.Println("Containers running: ", item.Info.ContainersRunning)
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
