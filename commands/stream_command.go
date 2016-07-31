package commands

import "github.com/tucnak/telebot"

type StreamCommandState string

const (
	InitializeStreamCommandState = "initalize"
)

type StreamCommand struct {
	*BasicCommand
	Telegram *telebot.Bot
	state    StreamCommandState
	qitems []string
}

type StreamNewCommand struct {
	*StreamCommand
}

func NewStreamCommand(telegram *telebot.Bot) *StreamCommand {
	cmd := StreamCommand{Telegram: telegram, state: InitializeStreamCommandState}
	return &cmd
}

func NewStreamNewCommand(telegram *telebot.Bot) StreamNewCommand {
	cmd := StreamNewCommand{NewStreamCommand(telegram)}
	return cmd
}

func (cmd StreamCommand) Execute(message telebot.Message) Commander {
	return nil
}