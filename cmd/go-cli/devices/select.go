package devices

import (
	log "github.com/sirupsen/logrus"

	"github.com/AlecAivazis/survey/v2"
	"github.com/docker/docker/api/types"
)

func getLocalDeviceAddress(index int, deviceInfo []types.Info) string {
	return deviceInfo[index].Name + ".local"
}

func selectLocalDevice() (int, []types.Info) {
	var device int
	deviceInfo := arpScan()

	if len(deviceInfo) == 0 {
		return device, deviceInfo
	}

	result := make([]string, len(deviceInfo))
	for i, item := range deviceInfo {
		result[i] = item.Name
	}

	prompt := &survey.Select{
		Message: "Select a device:",
		Options: result,
	}

	err := survey.AskOne(prompt, &device)

	if err != nil {
		log.Fatal(err)
	}

	return device, deviceInfo
}
