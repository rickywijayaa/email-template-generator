package main

import (
	"email-template-generator/database"
	log "email-template-generator/log"
	"email-template-generator/routes"
	"fmt"
	"os"
)

func main() {
	_ = database.New()
	server := routes.New()

	log.Info("Hello , 世界")

	port := os.Getenv("RUNNING_PORT")
	if port == "" {
		port = "9090"
	}

	fmt.Println(port)

	server.Run(port)
}
