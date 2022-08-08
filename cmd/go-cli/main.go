package main

import (
	"flag"

	"github.com/balena-community/go-cli/cmd/go-cli/api"
	"github.com/balena-community/go-cli/cmd/go-cli/devices"
)

var (
	scan     bool
	startApi bool
)

func init() {
	// Set flags
	flag.BoolVar(&scan, "scan", false, "Scan for balenaOS development devices on your local network")
	flag.BoolVar(&startApi, "api", false, "Start the API and listen for requests")
}

func main() {
	// Parse all flags from all files
	flag.Parse()

	// Take action based on flag
	switch {
	case startApi:
		api.Start()
	case scan:
		devices.Scan()
	default:
		flag.PrintDefaults()
	}
}
