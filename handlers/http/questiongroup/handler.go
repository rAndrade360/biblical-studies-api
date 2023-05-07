package questiongroup

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	dto "github.com/rAndrade360/biblical-studies-api/dto/questiongroup"
	"github.com/rAndrade360/biblical-studies-api/internal/models"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
	"github.com/rAndrade360/biblical-studies-api/services/questiongroup"
)

type controller struct {
	service questiongroup.QuestionGroupService
}

type QuestionGroupController interface {
	Create(ctx *fiber.Ctx) error
	List(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
}

func NewQuestionGroupController(svc questiongroup.QuestionGroupService) QuestionGroupController {
	return &controller{
		service: svc,
	}
}

func (c *controller) Create(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	log := ctx.Locals(logger.LogKey).(logger.Logger)
	contxt := context.WithValue(context.Background(), logger.LogKey, log)

	var in dto.QuestionGroupHttpCreate
	err := ctx.BodyParser(&in)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	qg := models.QuestionGroup{
		Name:        in.Name,
		Description: in.Description,
		ImageUrl:    in.ImageUrl,
		PrevQGID:    in.PrevQGID,
	}

	err = c.service.Create(contxt, &qg)
	if err != nil {
		log.Error(err.Error())
		return ctx.Status(500).Send([]byte(`{"message": "internal server error"}`))
	}

	jb, err := json.Marshal(dto.QuestionGroupHttpCreateResponse(qg))
	if err != nil {
		log.Error(err.Error())
		return ctx.Status(500).Send([]byte(`{"message": "internal server error"}`))
	}

	log.Infof("response: %s", string(jb))

	return ctx.Status(201).Send(jb)
}

func (c *controller) GetById(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	log := ctx.Locals(logger.LogKey).(logger.Logger)
	contxt := context.WithValue(context.Background(), logger.LogKey, log)

	id := ctx.Params("id")

	qg, err := c.service.GetById(contxt, id)
	if err != nil {
		log.Error(err.Error())
		return ctx.Status(500).Send([]byte(`{"message": "internal server error"}`))
	}

	jb, err := json.Marshal(dto.QuestionGroupHttpCreateResponse(*qg))
	if err != nil {
		log.Error(err.Error())
		return ctx.Status(500).Send([]byte(`{"message": "internal server error"}`))
	}
	log.Infof("response: %s", string(jb))

	return ctx.Status(200).Send(jb)
}

func (c *controller) List(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")
	log := ctx.Locals(logger.LogKey).(logger.Logger)
	contxt := context.WithValue(context.Background(), logger.LogKey, log)

	qgs, err := c.service.List(contxt)
	if err != nil {
		log.Error(err.Error())
		return ctx.Status(500).Send([]byte(`{"message": "internal server error"}`))
	}

	var res []dto.QuestionGroupHttpCreateResponse

	for i := range qgs {
		res = append(res, dto.QuestionGroupHttpCreateResponse(qgs[i]))
	}

	jb, err := json.Marshal(res)
	if err != nil {
		log.Error(err.Error())
		return ctx.Status(500).Send([]byte(`{"message": "internal server error"}`))
	}
	log.Infof("response: %s", string(jb))

	return ctx.Status(200).Send(jb)
}
