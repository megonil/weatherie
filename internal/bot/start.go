// Package bot
package bot

import (
	"weatherie/initializers"
	"weatherie/internal/utils"

	"github.com/go-telegram/bot"
	"log"
)

func Start() {
	opts := []bot.Option{
		bot.WithDefaultHandler(currentWeatherHandler),
		bot.WithCheckInitTimeout(utils.Seconds(10)),
	}

	b, err := bot.New(initializers.BotToken, opts...)
	if err != nil {
		log.Fatalf("Fatal error occured when initing bot: %v\n", err)
	}

	log.Println("Bot started")

	b.Start(initializers.Ctx)
}
