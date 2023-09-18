package main

import (
	"GetCryptoBot/bot"
	"GetCryptoBot/crypto"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	fmt.Println("Starting the program...")
	// Load environment variables from .env file using the bot package
	if err := bot.LoadConfig(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	pref := tele.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		// Get the user's first name
		firstName := c.Sender().FirstName

		// Create the welcome message
		welcomeMessage := fmt.Sprintf("Hi, %s, nice to meet you!\n"+
			"Enter the cryptocurrency whose official exchange rate\n"+
			"you want to know in relation to USD.\n"+
			"For example: BTC", firstName)

		// Send the welcome message
		return c.Send(welcomeMessage)
	})

	b.Handle(tele.OnText, func(c tele.Context) error {

		cryptoSymbol := strings.ToUpper(c.Text())
		result, err := crypto.GetCurrencyRate(cryptoSymbol)
		if err != nil {
			log.Fatal(err)
			return err
		}
		// Check if the result is empty or contains an error message
		if result == "0.000000000000000" || strings.Contains(result, "error") {
			return c.Send("Invalid or unsupported cryptocurrency symbol. Please enter a valid symbol.")
		}
		return c.Send(result)
	})

	b.Start()
}
