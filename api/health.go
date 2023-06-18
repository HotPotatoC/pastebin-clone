package api

import "github.com/gofiber/fiber/v2"

func (d *Dependency) Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "healthy",
	})
}
