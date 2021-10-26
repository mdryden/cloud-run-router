package main

import (
	"cloud-run-router/settings"
	"fmt"
	"log"
	"os"
)

func main() {

	app := App{}

	config := settings.GetConfig()

	app.Initialize(config)

	port, exists := os.LookupEnv("PORT")
	log.Print("Starting up...")

	if !exists {
		port = "8000"
	}

	port = fmt.Sprintf(":%s", port)

	app.Run(port)
}
