package questiongroup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	dto "github.com/rAndrade360/biblical-studies-api/dto/questiongroup"
	"github.com/rAndrade360/biblical-studies-api/internal/models"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
	"github.com/rAndrade360/biblical-studies-api/services/questiongroup"
)

type controller struct {
	service questiongroup.QuestionGroupService
	log     logger.Logger
}

type QuestionGroupController interface {
	Create(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
}

func NewQuestionGroupController(svc questiongroup.QuestionGroupService, logger logger.Logger) QuestionGroupController {
	return &controller{
		service: svc,
		log:     logger,
	}
}

func (c *controller) Create(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	c.log = c.log.SetRequestID(uuid.NewString())

	var in dto.QuestionGroupHttpCreate
	err := ctx.BodyParser(&in)
	if err != nil {
		c.log.Error(err.Error())
		return err
	}

	qg := models.QuestionGroup{
		Name:        in.Name,
		Description: in.Description,
		ImageUrl:    in.ImageUrl,
		PrevQGID:    in.PrevQGID,
	}

	err = c.service.Create(&qg)
	if err != nil {
		c.log.Error(err.Error())
		return ctx.Status(500).Send([]byte(`{"message": "internal server error"}`))
	}

	return ctx.Status(201).JSON(dto.QuestionGroupHttpCreateResponse(qg))
}

func (c *controller) GetById(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	c.log = c.log.SetRequestID(uuid.NewString())

	id := ctx.Params("id")

	qg, err := c.service.GetById(id)
	if err != nil {
		c.log.Error(err.Error())
		return ctx.Status(500).Send([]byte(`{"message": "internal server error"}`))
	}

	return ctx.Status(200).JSON(dto.QuestionGroupHttpCreateResponse(*qg))
}
