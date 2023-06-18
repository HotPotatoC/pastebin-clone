package api

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Error(c *fiber.Ctx, err error) error {
	if err.Error() == "Method Not Allowed" {
		return c.Status(http.StatusMethodNotAllowed).JSON(fiber.Map{
			"message": "Method not allowed",
		})
	}

	log.Err(err).Send()

	if os.Getenv("ENVIRONMENT") == "dev" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.
		Status(http.StatusInternalServerError).
		JSON(fiber.Map{
			"message": "An error has occurred on our side",
		})
}
