package mellody_bot

import (
	"github.com/tucnak/telebot"
	"sync"
	"time"
	"log"
	"fmt"
	"math/rand"
	"github.com/avatar29A/mellody-bot/commands"
)

type TelegramBot struct {
	ApiKey          string
	Telegram        *telebot.Bot
	globalWaitGroup *sync.WaitGroup
}

type WorkerInputChannel chan telebot.Message
type WorkerOutputChannel chan string

type Worker struct {
	UserName  string
	ChannelIn WorkerInputChannel
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

func dispatcher(messagesChannel chan telebot.Message, bot TelegramBot) {
	defer bot.globalWaitGroup.Done()
	workers := make(map[string]Worker)

	for message := range messagesChannel {
		worker, isWorkerExists := workers[message.Chat.Username]
		if isWorkerExists {
			worker.ChannelIn <- message
		} else {
			worker := Worker{UserName: message.Chat.Username, ChannelIn: make(WorkerInputChannel)}
			workers[message.Chat.Username] = worker

			go commandHandler(worker)
			worker.ChannelIn <- message
		}
	}
}

func commandHandler(worker Worker) {
	log.Print(fmt.Sprintf("Run worker for user: %v\n", worker.UserName))
	id := rand.Int()

	for {
		message := <-worker.ChannelIn
		log.Print(fmt.Sprintf("[Thread %v] Got message from %v: %v\n", id, worker.UserName, message.Text))
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