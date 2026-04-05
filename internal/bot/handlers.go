package bot

import (
	"context"
	"weatherie/internal/weather"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}

func currentWeatherHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	r := weather.Current(update.Message.Text)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   r,
	})
}
