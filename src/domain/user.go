package domain

import (
	"context"
	"errors"
	"soat23-gp14-hackaton-serverless/models"
	"soat23-gp14-hackaton-serverless/services"
)

type Users struct {
	provider services.Auth
}

var (
	ErrInvalidRegistry = errors.New("registry informed is invalid")
)

func NewUsersDomain(p services.Auth) *Users {
	return &Users{
		provider: p,
	}
}

func (u *Users) CreateUser(ctx context.Context, form models.UserForm) error {
	isValid := validateRegistry(form.Registry)
	if !isValid {
		return ErrInvalidRegistry
	}
	err := u.provider.SignUp(ctx, models.UserForm{
		Name:     form.Name,
		Registry: form.Registry,
		Password: form.Password,
		Email:    form.Email,
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *Users) Login(ctx context.Context, form models.UserLogin) (string, error) {
	isValid := validateRegistry(form.Username)
	if !isValid {
		return "", ErrInvalidRegistry
	}
	accessToken, err := u.provider.Login(ctx, form)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func validateRegistry(registry string) bool {
	if len(registry) != 5 {
		return false
	}

	return true
}
