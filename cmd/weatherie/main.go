package main

import (
	"log"
	"weatherie/initializers"
	"weatherie/internal/bot"
)

func init() {
	err := initializers.InitializeSecrets()
	if err != nil {
		log.Fatalf("Error during initializing: %v\n", err)
	}

	err = initializers.EnsureRedisConnected()
	if err != nil {
		log.Fatalf("Error connecting to the Redis: %e\n", err)
	}
}

func main() {
	defer initializers.Cancel()
	defer initializers.Rdb.Close()
	bot.Start()
}
