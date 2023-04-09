package questiongroup

import (
	"log"

	"github.com/gofiber/fiber/v2"
	dto "github.com/rAndrade360/biblical-studies-api/dto/questiongroup"
	"github.com/rAndrade360/biblical-studies-api/internal/models"
	"github.com/rAndrade360/biblical-studies-api/services/questiongroup"
)

type controller struct {
	service questiongroup.QuestionGroupService
}

type QuestionGroupController interface {
	Create(ctx *fiber.Ctx) error
}

func NewQuestionGroupController(svc questiongroup.QuestionGroupService) QuestionGroupController {
	return &controller{
		service: svc,
	}
}

func (c *controller) Create(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "application/json")

	var in dto.QuestionGroupHttpCreate
	err := ctx.BodyParser(&in)
	if err != nil {
		log.Println(err.Error())
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
		log.Println(err.Error())
		return ctx.Status(500).Send([]byte(`{"message": "internal server error"}`))
	}

	return ctx.Status(201).JSON(dto.QuestionGroupHttpCreateResponse(qg))
}
