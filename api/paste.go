package api

import (
	"errors"

	"github.com/HotPotatoC/pastebin-clone/backend"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
)

func (d Dependency) SavePaste(c *fiber.Ctx) error {
	body := c.Body()

	shortLink, err := d.Backend.SavePaste(c.Context(), backend.SavePasteParams{
		Text:          body,
		UserID:        c.Locals("userID").(string),
		UserIPAddress: c.IP(),
	})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).Send([]byte(BaseURL() + "/" + shortLink))
}

func (d Dependency) GetPaste(c *fiber.Ctx) error {
	shortLink := c.Params("short_link")

	paste, err := d.Backend.GetPaste(c.Context(), shortLink)
	if err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Paste not found",
			})
		}
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(paste)
}
