package devices

import (
	"fmt"
	"os"
	"os/exec"
)

func Ssh(address string, username string, port string) {
	if port == "" {
		port = "22222"
	}

	if username == "" {
		username = "root"
	}

	cmd := exec.Command("ssh", username+"@"+address, "-p", port)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("error connecting to host")
	}
}
