package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	apicontrollers "github.com/rAndrade360/biblical-studies-api/api/handlers/http"
	botcontroller "github.com/rAndrade360/biblical-studies-api/bot/handlers/http"
	"github.com/rAndrade360/biblical-studies-api/internal/infra/database/sqlite"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
	mwlogger "github.com/rAndrade360/biblical-studies-api/pkg/middlewares/logger"
)

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	TOKEN := os.Getenv("TOKEN")
	PORT := os.Getenv("PORT")

	if len(PORT) == 0 {
		PORT = "8080"
	}

	db, err := sqlite.New()
	if err != nil {
		log.Fatal("Err to connect db: ", err.Error())
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(mwlogger.Logger(logger.DEBUG))

	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Fatal("Err connect to bot: ", err.Error())
	}

	apiCtrls := apicontrollers.Load(db)
	botCtrls := botcontroller.Load(bot)

	app.Post("/bot", botCtrls.BotController.Handle)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).SendString(`{"message": "health is good"}`)
	})

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
