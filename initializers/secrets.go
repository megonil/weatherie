// Package initializers
package initializers

import (
	"os"

	"github.com/joho/godotenv"
)

func load(varn string) string {
	return os.Getenv(varn)
}

var BotToken string

func InitializeSecrets() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	BotToken = load("WEATHERIE_BOT_TOKEN")
	return nil
}
