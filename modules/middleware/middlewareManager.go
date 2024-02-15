package middleware

import (
	"context"
	"paradise-booking/config"
	"paradise-booking/entities"
)

type AccountStorage interface {
	GetAccountByEmail(ctx context.Context, email string) (*entities.Account, error)
}

type middlewareManager struct {
	cfg        *config.Config
	accountSto AccountStorage
}

func NewMiddlewareManager(cfg *config.Config, accountSto AccountStorage) *middlewareManager {
	return &middlewareManager{cfg: cfg, accountSto: accountSto}
}
