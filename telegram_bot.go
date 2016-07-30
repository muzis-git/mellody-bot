package mellody_bot

import (
	"github.com/tucnak/telebot"
	"sync"
	"time"
	"log"
)

type TelegramBot struct {
	ApiKey          string
	Telegram        *telebot.Bot
	globalWaitGroup *sync.WaitGroup
}

func NewTelegramBot(key string) TelegramBot {
	bot := TelegramBot{ApiKey: key, globalWaitGroup: &sync.WaitGroup{}}

	api, err := telebot.NewBot(bot.ApiKey)
	if err != nil {
		panic(err)
	}

	bot.Telegram = api
	return bot
}

func handler(messagesChannel chan telebot.Message, bot TelegramBot) {
	defer bot.globalWaitGroup.Done()

	for message := range messagesChannel {
		log.Printf("Received message '%v' from {%v}", message.Text, message.Chat.Username)

		switch message.Text {
		case "/hi":
			bot.Telegram.SendMessage(message.Chat,
				"Hello, " + message.Sender.FirstName + "!", nil)
		case "/clip":
			bot.Telegram.SendMessage(message.Chat, "https://www.youtube.com/watch?v=86URGgqONvA", nil)
		case "/exit":
			return
		default:
			bot.Telegram.SendMessage(message.Chat, "Try: /hi or /clip commands", nil)
		}
	}
}

func (bot TelegramBot) Start() {
	messagesChannel := make(chan telebot.Message)
	bot.Telegram.Listen(messagesChannel, 1 * time.Second)

	bot.globalWaitGroup.Add(1)

	go handler(messagesChannel, bot)

	bot.globalWaitGroup.Wait()
}

func (bot TelegramBot) Stop() {
	bot.globalWaitGroup.Done()
}