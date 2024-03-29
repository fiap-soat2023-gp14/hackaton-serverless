package services

import (
	"context"

	"soat23-gp14-hackaton-serverless/models"
)

type Auth interface {
	SignUp(context.Context, models.UserForm) error
	Login(context.Context, models.UserLogin) (string, error)
}
