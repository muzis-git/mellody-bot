package commands

import (
	"github.com/tucnak/telebot"
	"github.com/avatar29A/mellody-bot/utils"
	"log"
)

type StartCommand struct {
	*BasicCommand
	Telegram *telebot.Bot
}

func NewStartCommand(telegram *telebot.Bot) StartCommand {
	cmd := StartCommand{Telegram: telegram}
	return cmd
}

func (cmd StartCommand) Execute(message telebot.Message) {

	values := make(map[string]string)
	values["{{username}}"] = message.Sender.FirstName

	response, err := utils.LoadTemplate(utils.WELCOME_TEMPLATE_NAME, values)
	if err != nil {
		log.Printf("(500) StartCommand.Execute() crach after tried to load template.\n%v\n", err)
		return
	}

	cmd.Telegram.SendMessage(message.Sender,
		response,
		&telebot.SendOptions{
			ParseMode: telebot.ModeMarkdown,
		})
}

