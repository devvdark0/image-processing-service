package repository

import (
	"context"

	"github.com/devvdark0/image-processing-service/services/auth/internal/model"
)

type Repository interface {
	Create(ctx context.Context, email, password string) (int64, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}
