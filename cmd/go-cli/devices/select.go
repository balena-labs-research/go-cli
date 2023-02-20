package devices

import (
	"github.com/balena-labs-research/go-cli/pkg/networking"
	log "github.com/sirupsen/logrus"

	"github.com/AlecAivazis/survey/v2"
)

func getLocalDeviceAddress(index int, deviceInfo []networking.DockerResponse) string {
	return deviceInfo[index].Address
}

func selectLocalDevice() (int, []networking.DockerResponse) {
	var device int
	deviceInfo := lookupScan("")

	if len(deviceInfo) == 0 {
		return device, deviceInfo
	}

	result := make([]string, len(deviceInfo))
	for i, item := range deviceInfo {
		result[i] = item.Info.Name
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
