package mellody_bot

import (
	"github.com/tucnak/telebot"
	"sync"
	"time"
	"log"
	"fmt"
	"github.com/avatar29A/mellody-bot/commands"
)

type TelegramBot struct {
	apiKey          string
	Telegram        *telebot.Bot
	globalWaitGroup *sync.WaitGroup
}

type WorkerInputChannel chan telebot.Message
type WorkerOutputChannel chan string

type Worker struct {
	UserName  string
	ChannelIn WorkerInputChannel
	Bot       TelegramBot
}

func NewTelegramBot(key string) TelegramBot {
	bot := TelegramBot{apiKey: key, globalWaitGroup: &sync.WaitGroup{}}

	api, err := telebot.NewBot(bot.apiKey)
	if err != nil {
		panic(err)
	}

	bot.Telegram = api
	return bot
}

func dispatcher(messagesChannel chan telebot.Message, bot TelegramBot) {
	defer bot.globalWaitGroup.Done()
	workers := make(map[string]Worker)

	for message := range messagesChannel {
		worker, isWorkerExists := workers[message.Chat.Username]
		if isWorkerExists {
			worker.ChannelIn <- message
		} else {
			worker := Worker{UserName: message.Chat.Username, Bot: bot, ChannelIn: make(WorkerInputChannel)}
			workers[message.Chat.Username] = worker

			go commandHandler(worker)
			worker.ChannelIn <- message
		}
	}
}

func commandHandler(worker Worker) {
	log.Print(fmt.Sprintf("Run worker for user: %v\n", worker.UserName))
	//var currentCommand commands.Commander = nil

	for {
		message := <-worker.ChannelIn
		if message.Text == "/start" {
			cmd := commands.NewStartCommand(worker.Bot.Telegram)
			cmd.Execute(message)
		}
	}
}

func (bot TelegramBot) Start() {
	messagesChannel := make(chan telebot.Message)
	bot.Telegram.Listen(messagesChannel, 1 * time.Second)

	bot.globalWaitGroup.Add(1)

	go dispatcher(messagesChannel, bot)

	bot.globalWaitGroup.Wait()
}

func (bot TelegramBot) Stop() {
	bot.globalWaitGroup.Done()
}