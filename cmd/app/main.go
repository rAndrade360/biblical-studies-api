package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	altcontroller "github.com/rAndrade360/biblical-studies-api/handlers/http/alternative"
	qcontroller "github.com/rAndrade360/biblical-studies-api/handlers/http/question"
	qgcontroller "github.com/rAndrade360/biblical-studies-api/handlers/http/questiongroup"
	"github.com/rAndrade360/biblical-studies-api/internal/infra/database/sqlite"
	alrepository "github.com/rAndrade360/biblical-studies-api/internal/repositories/alternative"
	qrepository "github.com/rAndrade360/biblical-studies-api/internal/repositories/question"
	qgrepository "github.com/rAndrade360/biblical-studies-api/internal/repositories/questiongroup"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
	mwlogger "github.com/rAndrade360/biblical-studies-api/pkg/middlewares/logger"
	altservice "github.com/rAndrade360/biblical-studies-api/services/alternative"
	qservice "github.com/rAndrade360/biblical-studies-api/services/question"
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

	app := fiber.New()
	app.Use(mwlogger.Logger(logger.DEBUG))

	qgrepo := qgrepository.NewQuestionGroupRepository(db)
	qrepo := qrepository.NewQuestionRepository(db)
	altrepo := alrepository.NewAlternativeRepository(db)

	qgsvc := qgservice.NewQuestionGroupService(qgrepo)
	qsvc := qservice.NewQuestionService(qrepo, qgsvc)
	altsvc := altservice.NewAlternativeService(altrepo, qsvc)

	qgctrl := qgcontroller.NewQuestionGroupController(qgsvc)
	qctrl := qcontroller.NewQuestionController(qsvc)
	altctrl := altcontroller.NewAlternativeController(altsvc)

	qgrouter := app.Group("/questiongroup")
	qgrouter.Post("/", qgctrl.Create)
	qgrouter.Get("/", qgctrl.List)
	qgrouter.Get("/:id", qgctrl.GetById)

	qrouter := app.Group("/question")
	qrouter.Post("/", qctrl.Create)
	qrouter.Get("/", qctrl.List)
	qrouter.Get("/:id", qctrl.GetById)
	qrouter.Get("/:id/alternatives", altctrl.GetByQuestionId)
	qrouter.Post("/:id/alternatives", altctrl.Create)

	log.Fatal(app.Listen(":" + PORT))
}
