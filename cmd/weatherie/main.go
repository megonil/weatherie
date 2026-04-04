package main

import (
	"log"
	"weatherie/initializers"
	"weatherie/internal/bot"
)

func init() {
	err := initializers.InitializeSecrets()
	if err != nil {
		log.Fatalf("Error during initializing: %v", err)
	}
}

func main() {
	bot.Start()
}
