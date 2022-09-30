package devices

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func Ssh(args []string, username string, port string) {
	var address string

	if port == "" {
		port = "22222"
	}

	if username == "" {
		username = "root"
	}

	if len(args) > 0 {
		address = args[0]
	} else {
		index, deviceInfo := selectLocalDevice()

		if len(deviceInfo) == 0 {
			fmt.Println("No devices found")
			return
		}

		address = getLocalDeviceAddress(index, deviceInfo)
	}

	cmd := exec.Command("ssh", username+"@"+address, "-p", port)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
