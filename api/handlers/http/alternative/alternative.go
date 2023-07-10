package alternative

import (
	"context"

	stderr "errors"

	"github.com/gofiber/fiber/v2"
	dto "github.com/rAndrade360/biblical-studies-api/api/dto/alternative"
	"github.com/rAndrade360/biblical-studies-api/api/internal/models"
	"github.com/rAndrade360/biblical-studies-api/api/services/alternative"
	errors "github.com/rAndrade360/biblical-studies-api/pkg/error"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
)

type controller struct {
	service alternative.AlternativeService
}

type AlternativeController interface {
	Create(ctx *fiber.Ctx) error
	List(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	GetByQuestionId(ctx *fiber.Ctx) error
}

func NewAlternativeController(svc alternative.AlternativeService) AlternativeController {
	return &controller{
		service: svc,
	}
}

func (c *controller) Create(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	log := ctx.Locals(logger.LogKey).(logger.Logger)
	contxt := context.WithValue(context.Background(), logger.LogKey, log)

	var in dto.AlternativeHttpRequest
	err := ctx.BodyParser(&in)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	questionId := ctx.Params("id")

	a := models.NewAlternative(questionId, in.Value, in.IsCorret)

	err = c.service.Create(contxt, &a)
	if err != nil {
		log.Error(err.Error())
		if stderr.Is(err, errors.INVALIDINPUT) {
			return ctx.Status(400).Send([]byte(errors.BAD_REQUEST_HTTP.Error()))
		}
		return ctx.Status(500).Send([]byte(errors.ITERNAL_SERVER_ERROR_HTTP.Error()))
	}

	res := dto.AlternativeHttpResponse{
		ID:         a.ID,
		Value:      a.Value,
		QuestionID: a.QuestionID,
	}

	log.Info("response", res)

	return ctx.Status(201).JSON(res)
}

func (c *controller) GetById(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	log := ctx.Locals(logger.LogKey).(logger.Logger)
	contxt := context.WithValue(context.Background(), logger.LogKey, log)

	id := ctx.Params("id")

	a, err := c.service.GetById(contxt, id)
	if err != nil {
		log.Error(err.Error())
		if stderr.Is(err, errors.INVALIDINPUT) {
			return ctx.Status(400).Send([]byte(errors.BAD_REQUEST_HTTP.Error()))
		}
		return ctx.Status(500).Send([]byte(errors.ITERNAL_SERVER_ERROR_HTTP.Error()))
	}

	res := dto.AlternativeHttpResponse{
		ID:         a.ID,
		Value:      a.Value,
		QuestionID: a.QuestionID,
	}

	log.Info("response", res)

	return ctx.Status(200).JSON(res)
}

func (c *controller) GetByQuestionId(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	log := ctx.Locals(logger.LogKey).(logger.Logger)
	contxt := context.WithValue(context.Background(), logger.LogKey, log)

	id := ctx.Params("id")

	as, err := c.service.GetByQuestionId(contxt, id)
	if err != nil {
		log.Error(err.Error())
		if stderr.Is(err, errors.INVALIDINPUT) {
			return ctx.Status(400).Send([]byte(errors.BAD_REQUEST_HTTP.Error()))
		}
		return ctx.Status(500).Send([]byte(errors.ITERNAL_SERVER_ERROR_HTTP.Error()))
	}

	var res []dto.AlternativeHttpResponse

	for i := range as {
		res = append(res, dto.AlternativeHttpResponse{
			ID:         as[i].ID,
			Value:      as[i].Value,
			QuestionID: as[i].QuestionID,
		})
	}

	log.Info("response", res)

	return ctx.Status(200).JSON(res)
}

func (c *controller) List(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	log := ctx.Locals(logger.LogKey).(logger.Logger)
	contxt := context.WithValue(context.Background(), logger.LogKey, log)

	as, err := c.service.List(contxt)
	if err != nil {
		log.Error(err.Error())
		if stderr.Is(err, errors.INVALIDINPUT) {
			return ctx.Status(400).Send([]byte(errors.BAD_REQUEST_HTTP.Error()))
		}
		return ctx.Status(500).Send([]byte(errors.ITERNAL_SERVER_ERROR_HTTP.Error()))
	}

	var res []dto.AlternativeHttpResponse

	for i := range as {
		res = append(res, dto.AlternativeHttpResponse{
			ID:         as[i].ID,
			Value:      as[i].Value,
			QuestionID: as[i].QuestionID,
		})
	}

	log.Info("response", res)

	return ctx.Status(200).JSON(res)
}
