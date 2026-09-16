// Command api starts the remote_flow backend HTTP server.
package main

import (
	"log"

	"github.com/jaydenyip/remote_flow_v2/backend/internal/config"
	"github.com/jaydenyip/remote_flow_v2/backend/internal/server"
)

func main() {
	configs := config.Load()
	server := server.New(configs)

	log.Print("Hello World")

	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}

}
