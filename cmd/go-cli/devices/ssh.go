package devices

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Ssh(args []string, username string, port string) {
	var address string
	var err error

	if port == "" {
		port = "22222"
	}

	if username == "" {
		username = "root"
	}

	if len(args) > 0 {
		address = args[0]
	} else {
		index, deviceInfo, err := selectLocalDevice()

		if err != nil {
			log.Printf("interface %v \n", err)
			log.Print("Check you are running as root")
			os.Exit(1)
		}

		if len(deviceInfo) == 0 {
			fmt.Println("No devices found")
			return
		}

		address = deviceInfo[index].Name + ".local"
	}

	cmd := exec.Command("ssh", username+"@"+address, "-p", port)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("error connecting to host")
	}
}
