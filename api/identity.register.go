package api

import (
	"errors"
	"net/mail"
	"time"

	"github.com/HotPotatoC/pastebin-clone/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (d Dependency) Register(c *fiber.Ctx) error {
	var input RegisterInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	if err := input.Validate(); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	identity := repository.UsersStruct{
		Id:        uuid.New().String(),
		CreatedAt: time.Now(),
		Email:     input.Email,
		Name:      input.Name,
		Password:  input.Password,
	}

	accessToken, err := d.Backend.Register(c.Context(), identity)
	if err != nil {
		if errors.Is(err, repository.ErrEmailAlreadyExists) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "Email already exists",
			})
		}
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":      "User registered successfully",
		"access_token": accessToken,
	})
}

func (i RegisterInput) Validate() error {
	if i.Name == "" {
		return errors.New("name is required")
	}

	if i.Email == "" {
		return errors.New("email is required")
	}

	if _, err := mail.ParseAddress(i.Email); err != nil {
		return errors.New("invalid email")
	}

	if i.Password == "" {
		return errors.New("password is required")
	}

	if len(i.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	return nil
}
