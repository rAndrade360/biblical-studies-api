package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	apicontrollers "github.com/rAndrade360/biblical-studies-api/api/handlers/http"
	"github.com/rAndrade360/biblical-studies-api/internal/infra/database/sqlite"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
	mwlogger "github.com/rAndrade360/biblical-studies-api/pkg/middlewares/logger"
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

	app := fiber.New()
	app.Use(mwlogger.Logger(logger.DEBUG))

	apiCtrls := apicontrollers.Load(db)

	qgrouter := app.Group("/questiongroup")
	qgrouter.Post("/", apiCtrls.QuestionGroupController.Create)
	qgrouter.Get("/", apiCtrls.QuestionGroupController.List)
	qgrouter.Get("/:id", apiCtrls.QuestionGroupController.GetById)

	qrouter := app.Group("/question")
	qrouter.Post("/", apiCtrls.QuestionController.Create)
	qrouter.Get("/", apiCtrls.QuestionController.List)
	qrouter.Get("/:id", apiCtrls.QuestionController.GetById)
	qrouter.Get("/:id/alternatives", apiCtrls.AlternativeController.GetByQuestionId)
	qrouter.Post("/:id/alternatives", apiCtrls.QuestionGroupController.Create)

	log.Fatal(app.Listen(":" + PORT))
}
