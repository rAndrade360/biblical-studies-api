package logger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
)

func Logger(level logger.LOG_LEVEL) func(*fiber.Ctx) error {
	reqID := uuid.NewString()
	l := logger.NewLogger(level).SetRequestID(reqID)
	return func(c *fiber.Ctx) error {
		c.Locals(logger.LogKey, l)
		c.Set("Requestid", reqID)
		return c.Next()
	}
}
