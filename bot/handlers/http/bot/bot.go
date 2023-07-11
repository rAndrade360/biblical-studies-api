package bot

import (
	//	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"context"
	"net/http"

	stderr "errors"

	"github.com/rAndrade360/biblical-studies-api/bot/services/router"
	errors "github.com/rAndrade360/biblical-studies-api/pkg/error"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type controller struct {
	bot       *tgbotapi.BotAPI
	routerSvc router.RouterService
}

type BotController interface {
	Handle(c *fiber.Ctx) error
}

func NewBotController(bot *tgbotapi.BotAPI, routerSvc router.RouterService) BotController {
	return &controller{
		routerSvc: routerSvc,
		bot:       bot,
	}
}

func (c *controller) Handle(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	log := ctx.Locals(logger.LogKey).(logger.Logger)
	contxt := context.WithValue(context.Background(), logger.LogKey, log)

	log.Info("Request")

	var r http.Request

	fasthttpadaptor.ConvertRequest(ctx.Context(), &r, true)

	update, err := c.bot.HandleUpdate(&r)
	if err != nil {
		return err
	}

	err = c.routerSvc.Handle(contxt, *update.Message)
	if err != nil {
		log.Error(err.Error())
		if stderr.Is(err, errors.INVALIDINPUT) {
			return ctx.Status(400).Send([]byte(errors.BAD_REQUEST_HTTP.Error()))
		}
		return ctx.Status(500).Send([]byte(errors.ITERNAL_SERVER_ERROR_HTTP.Error()))
	}

	return ctx.SendStatus(204)
}
