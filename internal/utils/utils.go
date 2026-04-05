// Package utils
package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Seconds(seconds uint) time.Duration {
	return time.Second * time.Duration(seconds)
}

func Msg(ctx context.Context, b *bot.Bot, update *models.Update, text string) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      text,
		ParseMode: models.ParseModeMarkdown,
	})
}

func ToTGMonospace(val string) string {
	return fmt.Sprintf("```\n%s\n```", val)
}
