package repository

import (
	"context"
	"github.com/cepljxiongjun/hotwind/models"
)

type UserRepo interface {
	Create(ctx context.Context, p *models.User) (int64, error)
	Fetch(ctx context.Context, username string) (*models.User, error)
}