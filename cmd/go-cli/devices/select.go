package devices

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/docker/docker/api/types"
)

func selectLocalDevice() (int, []types.Info, error) {
	var device int
	deviceInfo, err := StartArpScan()

	if err != nil {
		return device, deviceInfo, err
	}

	if len(deviceInfo) == 0 {
		return device, deviceInfo, nil
	}

	result := make([]string, len(deviceInfo))
	for i, item := range deviceInfo {
		result[i] = item.Name
	}

	prompt := &survey.Select{
		Message: "Select a device:",
		Options: result,
	}

	err = survey.AskOne(prompt, &device)

	if err != nil {
		return device, deviceInfo, err
	}

	return device, deviceInfo, nil
}
