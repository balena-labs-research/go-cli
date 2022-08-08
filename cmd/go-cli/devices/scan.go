package devices

import (
	"fmt"
	"log"
	"os"

	"github.com/balena-community/go-cli/pkg/networking"
	"github.com/balena-community/go-cli/pkg/spinner"
)

func Scan() {
	s := spinner.StartNew("Scanning for local balenaOS devices...")
	deviceInfo, err := networking.GetDevices()
	s.Stop()

	if err != nil {
		log.Printf("interface %v \n", err)
		log.Print("Check you are running as root")
		os.Exit(1)
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
