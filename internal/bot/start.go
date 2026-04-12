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
		bot.WithCheckInitTimeout(utils.Seconds(10)),
		bot.WithInitialOffset(100),
	}

	b, err := bot.New(initializers.BotToken, opts...)
	if err != nil {
		log.Fatalf("Fatal error occured when initing bot: %v\n", err)
	}

	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"start",
		bot.MatchTypeCommandStartOnly,
		startCmd)
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"config",
		bot.MatchTypeCommandStartOnly,
		configCmd)
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"weather",
		bot.MatchTypeCommandStartOnly,
		weatherCmd)
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"location",
		bot.MatchTypeCommandStartOnly,
		locationCmd)

	log.Println("Bot started")

	b.Start(initializers.Ctx)
}
