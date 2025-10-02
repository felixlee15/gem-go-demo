package taskiface

import (
	"context"

	"go-demo/ent"
)

type Repository interface {
	Create(ctx context.Context, title string, userID string) (*ent.Task, error)
	Update(ctx context.Context, taskID string, completed bool) (*ent.Task, error)
	List(ctx context.Context) ([]*ent.Task, error)
}
