package main

import (
	"gcp-api/api"
	"log"
)

func main() {
	// Start HTTP server with API and print any error
	err := api.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
