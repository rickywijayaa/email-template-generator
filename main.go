package main

import (
	"email-template-generator/database"
	"email-template-generator/routes"
	"fmt"
	"os"
)

func main() {
	db := database.New()
	server := routes.New(db)

	port := os.Getenv("RUNNING_PORT")
	if port == "" {
		port = "9090"
	}

	server.Run(fmt.Sprintf("localhost:%v", port))
}
