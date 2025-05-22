package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "Cписок комманд: \n" +
				"- /time - Показать время до приезда друзей \n" +
				"- /bd - Показать когда у друзей дни рождения \n"
		case "time":
			targetDate := time.Date(2025, 6, 7, 0, 0, 0, 0, time.UTC)
			now := time.Now()

			diff := targetDate.Sub(now)
			days := math.Ceil(diff.Hours() / 24)
			if days > 4 {
				msg.Text = fmt.Sprintf("До приезда друзей осталось: \n%.0f дней", days)
			} else {
				msg.Text = fmt.Sprintf("До приезда друзей осталось: \n%.0f дня", days)
			}
		case "bd":
			msg.Text = "Дни рождения друзей: \n" +
				"Сеня - 28.03\n" +
				"Леша - 03.05\n" +
				"Даша - 06.06\n" +
				"Усман - 08.06\n" +
				"Дима - 20.10\n"
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
