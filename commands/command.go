package commands

import "github.com/tucnak/telebot"

type Commander interface {
	IsEnded() bool
	Execute(message telebot.Message) Commander
}

type BasicCommand struct {
	isEnded bool
}

func (cmd *BasicCommand) IsEnded() bool {
	return cmd.isEnded
}
