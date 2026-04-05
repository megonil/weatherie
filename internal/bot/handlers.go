package bot

import (
	"context"
	"log"
	"weatherie/internal/utils"
	"weatherie/internal/weather"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// respond with the same
	utils.Msg(ctx, b, update, update.Message.Text)
}

func currentWeatherHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	c, err := weather.Current(update.Message.Text)
	if err != nil {
		utils.Msg(ctx, b, update, "Oopsie, something went wrong!")
		log.Printf("Error occured in response to the current weather: %e\n", err)
		return
	}

	b.SendPhoto(ctx, &bot.SendPhotoParams{
		ChatID:    update.Message.Chat.ID,
		Caption:   utils.ToTGMonospace(c.String()),
		Photo:     &models.InputFileString{Data: c.IconURL()},
		ParseMode: models.ParseModeMarkdown,
	})
}
