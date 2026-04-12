// Package initializers
package initializers

import (
	"context"
	"os"
	"os/signal"
	"strconv"
)

var Ctx, Cancel = signal.NotifyContext(context.Background(), os.Interrupt)

func load(varn string) string {
	return os.Getenv(varn)
}

var (
	BotToken        string
	WeatherAPIToken string
	RedisURL        string
	RedisPass       string
	RedisDBNum      int
)

func InitializeSecrets() error {
	var err error
	BotToken = load("BOT_TOKEN")
	WeatherAPIToken = load("WEATHER_API_TOKEN")
	RedisURL = load("REDIS_URL")
	RedisPass = load("REDIS_PASSWORD")
	dbString := load("REDIS_DB")
	RedisDBNum, err = strconv.Atoi(dbString)
	if err != nil {
		return err
	}

	return nil
}
