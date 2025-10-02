package taskiface

import (
	"context"

	"go-demo/internal/pkg/graph/models"
)

type UseCase interface {
	Create(ctx context.Context, title string, userID string) (*models.Task, error)
	Update(ctx context.Context, taskID string, completed bool) (*models.Task, error)
	List(ctx context.Context) ([]*models.Task, error)
}
