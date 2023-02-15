package devices

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/balena-labs-research/go-cli/pkg/networking"
	"github.com/balena-labs-research/go-cli/pkg/spinner"
	"github.com/docker/docker/api/types"
)

func Scan(scanType string, args []string) {
	var err error
	deviceInfo := []types.Info{}
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
	for deviceNumber, info := range deviceInfo {
		fmt.Printf(" - Device %v - \n", deviceNumber+1)
		fmt.Printf("Hostname: %s.local \n", info.Name)
		fmt.Println("Containers running: ", info.ContainersRunning)
		fmt.Printf("Kernel Version: %s \n", info.KernelVersion)
		fmt.Printf("Operating System: %s \n", info.OperatingSystem)
		fmt.Printf("Architecture: %s \n", info.Architecture)
		fmt.Printf("Balena Engine Version: %s \n", info.ServerVersion)
		fmt.Printf(" \n")
	}
}

func arpScan() []types.Info {
	s := spinner.StartNew("Scanning for local balenaOS devices...")

	// Arp scan for available devices
	arpResults, err := networking.ArpScan()

	if err != nil {
		log.Fatal("Check you are running as root")
	}

	deviceInfo := networking.CheckIpPorts(arpResults, 22222)

	// Stop before error to avoid overlap of messages
	s.Stop()

	return deviceInfo
}

func lookupScan(ipRange string) ([]types.Info, error) {
	s := spinner.StartNew("Scanning for local balenaOS devices...")
	lookupResults := networking.LookupAddresses(ipRange)

	deviceInfo := networking.CheckHostnamePorts(lookupResults, 22222)

	// Stop before error to avoid overlap of messages
	s.Stop()

	return deviceInfo, nil
}
