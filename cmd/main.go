package main

import (
	"github.com/avatar29A/mellody-bot"
	"os"
	"sync"
)

const (
	TELEGRAM_BOT_API_KEY = "TELEGRAM_BOT_API_KEY"
)

func startTelegramBot(bot mellody_bot.TelegramBot, wg *sync.WaitGroup) {
	defer wg.Done()
	bot.Start()
}

func main() {
	mainWaitGroup := &sync.WaitGroup{}

	bot := mellody_bot.NewTelegramBot(os.Getenv(TELEGRAM_BOT_API_KEY))
	mainWaitGroup.Add(1)
	go startTelegramBot(bot, mainWaitGroup)

	mainWaitGroup.Wait()
}