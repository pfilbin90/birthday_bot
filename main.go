package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Define a map of birthdays
	birthdays := map[string]string{
		"Patrick":  "1990-06-03",
		"Karl":     "1990-04-08",
		"Andreas":  "1990-06-26",
		"David":    "1990-10-27",
		"Benjamin": "1990-11-27",
		"Diego":    "1990-04-24",
		"Cameron":  "1990-05-21",
		"Will":     "1990-07-14",
		"Obaid":    "1990-12-19",
		"Raza":     "1990-03-17",
		"Peter":    "1990-07-01",
	}

	// Get the current date
	currentTime := time.Now()
	//currentYear := currentTime.Year()

	// Create a new Telegram bot with your bot token
	bot, err := tgbotapi.NewBotAPI("YOUR_TELEGRAM_BOT_TOKEN")
	if err != nil {
		log.Panic(err)
	}

	// Check if today is someone's birthday
	for name, birthday := range birthdays {
		// Parse the birthday string to get the birth date
		birthDate, err := time.Parse("2006-01-02", birthday)
		if err != nil {
			fmt.Println("Error parsing birthday:", err)
			continue
		}

		// Check if the day and month match
		if birthDate.Day() == currentTime.Day() && birthDate.Month() == currentTime.Month() {
			// Calculate age
			// age := currentYear - birthDate.Year()
			// Output the "Happy Birthday" message with age
			birthdayMessage := fmt.Sprintf("Happy Birthday, %s!", name)

			// Send the message to a Telegram chat
			chatID := int64(YOUR_CHAT_ID) // Replace YOUR_CHAT_ID with the actual chat ID
			message := tgbotapi.NewMessage(chatID, birthdayMessage)
			_, err := bot.Send(message)
			if err != nil {
				log.Println("Error sending birthday message:", err)
			}
		}
	}
}
