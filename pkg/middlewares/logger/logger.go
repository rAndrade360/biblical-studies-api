package logger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
)

func Logger(level logger.LOG_LEVEL) func(*fiber.Ctx) error {
	l := logger.NewLogger(level).SetRequestID(uuid.NewString())
	return func(c *fiber.Ctx) error {
		c.Locals(logger.LogKey, l)
		return c.Next()
	}
}
