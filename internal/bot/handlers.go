package bot

import (
	"context"
	"log"
	"strings"
	"weatherie/internal/utils"
	"weatherie/internal/weather"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// respond with the same
	utils.Msg(ctx, b, update, update.Message.Text)
}

func startCmd(ctx context.Context, b *bot.Bot, update *models.Update) {
	utils.Msg(ctx, b, update, "Hello! To get weather info at <location>, use command /weather <location>.\nYou also can specify location using latitude and longitude, /weather <lat>,<lon>. For location you also can use ip address.")
	InitConfig(update.Message.From.ID)
}

func configCmd(ctx context.Context, b *bot.Bot, update *models.Update) {
	config := GetConfigSafe(update.Message.From.ID)

	text := config.String()

	utils.Msg(ctx, b, update, text)
}

func weatherCmd(ctx context.Context, b *bot.Bot, update *models.Update) {
	// trim first word (the command)
	location := strings.Split(update.Message.Text, " ")[1:]
	str := strings.Join(location, " ")

	if !utils.IsLocationSafe(str) {
		utils.Msg(ctx, b, update, "Nice try.")
		return
	}
	if len(str) == 0 {
		config := GetConfigSafe(update.Message.From.ID)
		if config.Location != "" {
			str = config.Location
		} else {
			utils.Msg(ctx, b, update, "You hadn't saved any location and hadn't provided any. Usage: /weather <location>")
		}
	}

	c, err := weather.Current(str)

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

func locationCmd(ctx context.Context, b *bot.Bot, update *models.Update) {
	location := strings.Split(update.Message.Text, " ")[1:]
	if len(location) == 0 {
		utils.Msg(ctx, b, update, "Usage: /location <location to be saved>")
		return
	}

	str := strings.Join(location, " ")
	if !utils.IsLocationSafe(str) {
		utils.Msg(ctx, b, update, "Nice try.")
		return
	}

	config := GetConfigSafe(update.Message.From.ID)
	config.Location = str
	config.SavedLocation = true

	SaveConfig(config, update.Message.From.ID)
}
