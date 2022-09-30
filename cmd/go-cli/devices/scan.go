package devices

import (
	"fmt"
	"log"

	"github.com/balena-community/go-cli/pkg/networking"
	"github.com/balena-community/go-cli/pkg/spinner"
	"github.com/docker/docker/api/types"
)

func Scan() {
	deviceInfo := GetBalenaDevices()

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

func GetBalenaDevices() []types.Info {
	s := spinner.StartNew("Scanning for local balenaOS devices...")
	deviceInfo, err := networking.ScanBalenaDevices()

	// Stop before error to avoid overlap of messages
	s.Stop()

	if err != nil {
		log.Fatal("Check you are running as root")
	}

	return deviceInfo
}
