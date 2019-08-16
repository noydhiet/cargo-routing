package main

import (
	"cargo-tracking/transport"
	"go-kit/kit/log"
	"net/http"
	"os"
)

func main() {

	//service.HelloDaerah(string, string, string)

	logger := log.NewLogfmtLogger(os.Stdout)

	transport.RegisterHttpsServicesAndStartListener()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	logger.Log("listening-on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log("listen.error", err)
	}
}
