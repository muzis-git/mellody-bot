package main

import (
	"fmt"
	"github.com/avatar29A/mellody-bot"
	"os"
)

const (
	TELEGRAM_BOT_API_KEY = "TELEGRAM_BOT_API_KEY"
)

func main() {
	bot := mellody_bot.NewTelegramBot(os.Getenv(TELEGRAM_BOT_API_KEY))
	bot.Start()
	fmt.Print("Hello hakatone!")
}