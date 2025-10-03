package taskiface

import (
	"context"

	"gemdemo/ent"
)

type Repository interface {
	Create(ctx context.Context, title string, userID uint64) (*ent.Task, error)
	Update(ctx context.Context, taskID uint64, completed bool) (*ent.Task, error)
	List(ctx context.Context) ([]*ent.Task, error)
	GetByID(ctx context.Context, id uint64) (*ent.Task, error)
	GetByUser(ctx context.Context, userID uint64, completed *bool) ([]*ent.Task, error)
	GetOwner(ctx context.Context, obj *ent.Task) (*ent.User, error)
}
