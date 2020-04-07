/*
 * ModelDB API
 *
 * No description provided
 *
 * API version: 0.0.1
 */

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/medienlampe/modb-server/home"
	"github.com/medienlampe/modb-server/server"
)

var (
	// ModBCertFile    = os.Getenv("MODB_CERT_FILE")
	// ModBKeyFile     = os.Getenv("MODB_KEY_FILE")
	// ModBServiceAddr = os.Getenv("MODB_SERVICE_ADDR")

	// ModBServiceAddr is the address/port to listen to
	ModBServiceAddr = ":8080"
)

func main() {
	logger := log.New(os.Stdout, "modb-server ", log.LstdFlags|log.Lshortfile)

	h := home.NewHandlers(logger)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(mux, ModBServiceAddr)

	logger.Println("Starting...")
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("Failed to start: %v", err)
	}
}
