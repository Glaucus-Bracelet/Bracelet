package main

import (
	_ "github.com/joho/godotenv/autoload" // Automatically loads godotenv .env file contents. 💖

	"github.com/Glaucus/Bracelet/httpServer"
)

func main() {
	// Start http server
	httpServer.Start()
}
