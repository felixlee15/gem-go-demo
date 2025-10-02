package useriface

import (
	"context"

	"go-demo/internal/pkg/graph/models"
)

type UseCase interface {
	Create(ctx context.Context, name string, email string) (*models.User, error)
	List(ctx context.Context) ([]*models.User, error)
}
