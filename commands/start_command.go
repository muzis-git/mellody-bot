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

func (cmd StartCommand) Execute(message telebot.Message) Commander {
	response, err := utils.LoadTemplate(utils.WELCOME_TEMPLATE_NAME, nil)
	if err != nil {
		log.Printf("(500) StartCommand.Execute() crach after tried to load template.\n%v\n", err)

	} else {
		cmd.Telegram.SendMessage(message.Sender,
			response,
			&telebot.SendOptions{
				ParseMode: telebot.ModeMarkdown,
			})
	}

	return nil
}

