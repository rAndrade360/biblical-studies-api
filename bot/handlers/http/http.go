package http

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	botctrl "github.com/rAndrade360/biblical-studies-api/bot/handlers/http/bot"
	"github.com/rAndrade360/biblical-studies-api/bot/services/router"
)

type HttpControllers struct {
	BotController botctrl.BotController
}

func Load(bot *tgbotapi.BotAPI) HttpControllers {
	routerSvc := router.NewRouterService(bot)

	botCtrl := botctrl.NewBotController(bot, routerSvc)
	return HttpControllers{
		BotController: botCtrl,
	}
}
