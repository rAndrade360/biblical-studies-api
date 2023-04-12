package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	qgcontroller "github.com/rAndrade360/biblical-studies-api/handlers/http/questiongroup"
	"github.com/rAndrade360/biblical-studies-api/internal/infra/database/sqlite"
	qgrepository "github.com/rAndrade360/biblical-studies-api/internal/repositories/questiongroup"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
	qgservice "github.com/rAndrade360/biblical-studies-api/services/questiongroup"
)

func main() {
	db, err := sqlite.New()
	if err != nil {
		log.Fatal("Err to connect db: ", err.Error())
	}

	logger := logger.NewLogger(logger.DEBUG)

	qgrepo := qgrepository.NewQuestionGroupRepository(db)
	qgsvc := qgservice.NewQuestionGroupService(qgrepo, logger)

	qgctrl := qgcontroller.NewQuestionGroupController(qgsvc, logger)

	app := fiber.New()

	app.Post("/questiongroup", qgctrl.Create)
	app.Get("/questiongroup/:id", qgctrl.GetById)
	app.Listen(":8080")
}
