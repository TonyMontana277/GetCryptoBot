package bot

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	// Import the path to your crypto service package
)

type TelegramBot struct {
	BotToken string
}

func NewTelegramBot() *TelegramBot {
	botToken := os.Getenv("BOT_TOKEN")
	return &TelegramBot{
		BotToken: botToken,
	}
}

func (tb *TelegramBot) Run() error {
	// Initialize Telegram bot
	bot, err := tgbotapi.NewBotAPI(tb.BotToken)
	if err != nil {
		return err
	}
	bot.Debug = true
	kreacher_currencyBot := "@kreacher_currencyBot"
	log.Printf("Authorized as %s", kreacher_currencyBot)

	// Handle updates and implement your bot's functionality here
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		messageText := update.Message.Text
		chatID := update.Message.Chat.ID

		if messageText == "/start" {
			startCommandReceived(chatID, update.Message.Chat.FirstName, bot)
		} else {
			// Handle other commands or messages here
		}
	}

	return nil
}

func startCommandReceived(chatID int64, name string, bot *tgbotapi.BotAPI) {
	answer := fmt.Sprintf("Hi, %s, nice to meet you!\n"+
		"Enter the cryptocurrency whose official exchange rate\n"+
		"you want to know in relation to USD.\n"+
		"For example: BTC", name)
	sendMessage(chatID, answer, bot)
}

func sendMessage(chatID int64, textToSend string, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(chatID, textToSend)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
