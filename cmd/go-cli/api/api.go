package api

import (
	"encoding/json"
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/balena-community/go-cli/pkg/networking"
)

var (
	port string
)

type errorMessage struct {
	Message string `json:"error"`
}

func init() {
	flag.StringVar(&port, "port", "7878", "Specify a port to run the API")
}

// Returns JSON array with info for each device found on network
func scan(w http.ResponseWriter, r *http.Request) {
	devices, err := networking.ScanBalenaDevices()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonResp, err := json.Marshal(errorMessage{err.Error()})
		if err != nil {
			log.Errorf("Error happened in JSON marshal. Err: %s", err)
			return
		}
		_, err = w.Write(jsonResp)
		if err != nil {
			log.Errorf("Error happened in writing response. Err: %s", err)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(devices)

	if err != nil {
		log.Errorf("Error happened in JSON marshal. Err: %s", err)
		return
	}
}

func Start() {
	// Routes
	http.HandleFunc("/v1/scan", scan)

	// Start API server
	log.Printf("API running on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error starting API server. Err: %s", err)
	}
}
