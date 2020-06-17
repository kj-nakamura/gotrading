package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"
)

type HealthCheck struct {
	Status int
	Result string
}

func main() {
	// utils.LoggingSettings(config.Config.LogFile)
	// controllers.StreamIngestionData()
	log.Println(StartWebServer())
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	ping := HealthCheck{http.StatusOK, "ok"}

	res, err := json.Marshal(ping)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func StartWebServer() error {
	http.HandleFunc("/health-check/", healthCheckHandler)

	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)
}
