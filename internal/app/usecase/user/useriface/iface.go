package useriface

import (
	"context"

	"gemdemo/ent"
)

type UseCase interface {
	Create(ctx context.Context, name string, email string) (*ent.User, error)
	CountCompleteTasks(ctx context.Context, obj *ent.User) (int, error)
	List(ctx context.Context) ([]*ent.User, error)
	ListUserTasks(ctx context.Context, obj *ent.User) ([]*ent.Task, error)
}
