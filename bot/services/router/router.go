package router

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type service struct {
	bot *tgbotapi.BotAPI
}

type RouterService interface {
	Handle(ctx context.Context, message tgbotapi.Message) error
}

func NewRouterService(bot *tgbotapi.BotAPI) RouterService {
	return &service{
		bot: bot,
	}
}

func (s *service) Handle(ctx context.Context, message tgbotapi.Message) error {
	if message.Command() == "start" {
		reply := tgbotapi.NewMessage(message.From.ID, "Eh nois")
		s.bot.Send(reply)
	}
	return nil
}
