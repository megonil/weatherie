// Package bot
package bot

import (
	"context"
	"log"
	"os"
	"os/signal"
	"weatherie/initializers"
	"weatherie/internal/utils"

	"github.com/go-telegram/bot"
)

func Start() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(DefaultHandler),
		bot.WithCheckInitTimeout(utils.Seconds(10)),
	}

	b, err := bot.New(initializers.BotToken, opts...)
	if err != nil {
		log.Fatalf("Fatal error occured when initing bot: %v", err)
	}

	log.Printf("Bot started")

	b.Start(ctx)
}
