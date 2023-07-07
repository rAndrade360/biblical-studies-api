package question

import (
	"context"

	stderr "errors"

	"github.com/gofiber/fiber/v2"
	dto "github.com/rAndrade360/biblical-studies-api/dto/question"
	"github.com/rAndrade360/biblical-studies-api/internal/models"
	errors "github.com/rAndrade360/biblical-studies-api/pkg/error"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
	"github.com/rAndrade360/biblical-studies-api/services/question"
)

type controller struct {
	service question.QuestionService
}

type QuestionController interface {
	Create(ctx *fiber.Ctx) error
	List(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
}

func NewQuestionController(svc question.QuestionService) QuestionController {
	return &controller{
		service: svc,
	}
}

func (c *controller) Create(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	log := ctx.Locals(logger.LogKey).(logger.Logger)
	contxt := context.WithValue(context.Background(), logger.LogKey, log)

	var in dto.QuestionHttpCreate
	err := ctx.BodyParser(&in)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	qg := models.NewQuestion(in.QuestionGroupID, in.Title, in.Description, in.BibleText, in.ImageUrl, in.SortNumber)

	err = c.service.Create(contxt, &qg)
	if err != nil {
		log.Error(err.Error())
		if stderr.Is(err, errors.INVALIDINPUT) {
			return ctx.Status(400).Send([]byte(errors.BAD_REQUEST_HTTP.Error()))
		}
		return ctx.Status(500).Send([]byte(errors.ITERNAL_SERVER_ERROR_HTTP.Error()))
	}

	res := dto.QuestionHttpCreateResponse(qg)

	log.Info("response", res)

	return ctx.Status(201).JSON(res)
}

func (c *controller) GetById(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	log := ctx.Locals(logger.LogKey).(logger.Logger)
	contxt := context.WithValue(context.Background(), logger.LogKey, log)

	id := ctx.Params("id")

	qg, err := c.service.GetById(contxt, id)
	if err != nil {
		log.Error(err.Error())
		if stderr.Is(err, errors.INVALIDINPUT) {
			return ctx.Status(400).Send([]byte(errors.BAD_REQUEST_HTTP.Error()))
		}
		return ctx.Status(500).Send([]byte(errors.ITERNAL_SERVER_ERROR_HTTP.Error()))
	}

	res := dto.QuestionHttpCreateResponse(*qg)

	log.Info("response", res)

	return ctx.Status(200).JSON(res)
}

func (c *controller) List(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	log := ctx.Locals(logger.LogKey).(logger.Logger)
	contxt := context.WithValue(context.Background(), logger.LogKey, log)

	qgs, err := c.service.List(contxt)
	if err != nil {
		log.Error(err.Error())
		if stderr.Is(err, errors.INVALIDINPUT) {
			return ctx.Status(400).Send([]byte(errors.BAD_REQUEST_HTTP.Error()))
		}
		return ctx.Status(500).Send([]byte(errors.ITERNAL_SERVER_ERROR_HTTP.Error()))
	}

	var res []dto.QuestionHttpCreateResponse

	for i := range qgs {
		res = append(res, dto.QuestionHttpCreateResponse(qgs[i]))
	}

	log.Info("response", res)

	return ctx.Status(200).JSON(res)
}
