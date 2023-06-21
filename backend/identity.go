package backend

import (
	"context"
	"errors"
	"time"

	"github.com/HotPotatoC/pastebin-clone/logic"
	"github.com/HotPotatoC/pastebin-clone/repository"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterParams struct {
	Email    string
	Name     string
	Password string
}

func (d Dependency) Register(ctx context.Context, params RegisterParams) (string, error) {
	// Check if email already exists
	_, err := d.Repository.GetUserByEmail(ctx, params.Email)
	if !errors.Is(err, gocql.ErrNotFound) {
		return "", repository.ErrEmailAlreadyExists
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	params.Password = string(passwordHash)

	userID := uuid.New().String()
	// Save user
	err = d.Repository.SaveUser(ctx, repository.User{
		Id:        userID,
		Name:      params.Name,
		Email:     params.Email,
		Password:  params.Password,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return "", err
	}

	accessToken, err := logic.GenerateJWT(map[string]any{
		"userID": userID,
		"email":  params.Email,
	})
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

type LoginOutput struct {
	AccessToken string `json:"access_token"`
	User        struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
	} `json:"user"`
}

func (d Dependency) Login(ctx context.Context, email, password string) (LoginOutput, error) {
	var output LoginOutput

	identity, err := d.Repository.GetUserByEmail(ctx, email)
	if err != nil {
		return LoginOutput{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(identity.Password), []byte(password)); err != nil {
		return LoginOutput{}, repository.ErrMismatchPassword
	}

	accessToken, err := logic.GenerateJWT(map[string]any{
		"userID": identity.Id,
		"email":  identity.Email,
	})
	if err != nil {
		return LoginOutput{}, repository.ErrMismatchPassword
	}

	output.AccessToken = accessToken
	output.User.Id = identity.Id
	output.User.Name = identity.Name
	output.User.Email = identity.Email
	output.User.CreatedAt = identity.CreatedAt.Format(time.RFC3339)

	return output, nil
}
