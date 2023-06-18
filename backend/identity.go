package backend

import (
	"context"
	"errors"

	"github.com/HotPotatoC/pastebin-clone/logic"
	"github.com/HotPotatoC/pastebin-clone/repository"
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
)

func (d Dependency) Register(ctx context.Context, identity repository.UsersStruct) (string, error) {
	// Check if email already exists
	_, err := d.Repository.GetUserByEmail(ctx, identity.Email)
	if !errors.Is(err, gocql.ErrNotFound) {
		return "", repository.ErrEmailAlreadyExists
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(identity.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	identity.Password = string(passwordHash)

	// Save user
	err = d.Repository.SaveUser(ctx, identity)
	if err != nil {
		return "", err
	}

	accessToken, err := logic.GenerateJWT(map[string]any{
		"userID": identity.Id,
		"email":  identity.Email,
	})
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (d Dependency) Login(ctx context.Context, email, password string) (string, repository.UsersStruct, error) {
	identity, err := d.Repository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", repository.UsersStruct{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(identity.Password), []byte(password)); err != nil {
		return "", repository.UsersStruct{}, repository.ErrMismatchPassword
	}

	accessToken, err := logic.GenerateJWT(map[string]any{
		"userID": identity.Id,
		"email":  identity.Email,
	})
	if err != nil {
		return "", repository.UsersStruct{}, repository.ErrMismatchPassword
	}

	return accessToken, identity, nil
}
