package api

import (
	"errors"
	"net/mail"

	"github.com/HotPotatoC/pastebin-clone/repository"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
)

type LoginInput struct {
	Email    string
	Password string
}

func (d *Dependency) Login(c *fiber.Ctx) error {
	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return Error(c, err)
	}

	if err := input.Validate(); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	output, err := d.Backend.Login(c.Context(), input.Email, input.Password)
	switch {
	case errors.Is(err, gocql.ErrNotFound):
		fallthrough
	case errors.Is(err, repository.ErrMismatchPassword):
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	case err != nil:
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "User logged in successfully",
		"access_token": output.AccessToken,
		"user":         output.User,
	})
}

func (i LoginInput) Validate() error {
	if i.Email == "" {
		return errors.New("email is required")
	}

	if _, err := mail.ParseAddress(i.Email); err != nil {
		return errors.New("invalid email")
	}

	if i.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
