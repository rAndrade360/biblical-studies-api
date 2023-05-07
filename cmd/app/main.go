package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	qgcontroller "github.com/rAndrade360/biblical-studies-api/handlers/http/questiongroup"
	"github.com/rAndrade360/biblical-studies-api/internal/infra/database/sqlite"
	qgrepository "github.com/rAndrade360/biblical-studies-api/internal/repositories/questiongroup"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
	mwlogger "github.com/rAndrade360/biblical-studies-api/pkg/middlewares/logger"
	qgservice "github.com/rAndrade360/biblical-studies-api/services/questiongroup"
)

var (
	PORT = os.Getenv("PORT")
)

func main() {
	if len(PORT) == 0 {
		PORT = "8080"
	}

	db, err := sqlite.New()
	if err != nil {
		log.Fatal("Err to connect db: ", err.Error())
	}

	//logger := logger.NewLogger(logger.DEBUG)

	qgrepo := qgrepository.NewQuestionGroupRepository(db)
	app := fiber.New()
	app.Use(mwlogger.Logger(logger.DEBUG))
	qgsvc := qgservice.NewQuestionGroupService(qgrepo)

	qgctrl := qgcontroller.NewQuestionGroupController(qgsvc)

	app.Post("/questiongroup", qgctrl.Create)
	app.Get("/questiongroup", qgctrl.List)
	app.Get("/questiongroup/:id", qgctrl.GetById)
	app.Listen(":" + PORT)
}
