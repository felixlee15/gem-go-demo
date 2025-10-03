package user

import (
	"context"
	"fmt"

	"gemdemo/db"
	"gemdemo/ent"
	"gemdemo/ent/task"
	"gemdemo/ent/user"
	"gemdemo/internal/app/repository/user/useriface"
)

type repo struct{}

func NewRepository() useriface.Repository {
	return &repo{}
}

func (instance *repo) Create(ctx context.Context, name string, email string) (*ent.User, error) {
	user, err := db.GetClient(ctx).User.Create().SetName(name).SetEmail(email).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("userRepo.Create: %w", err)
	}

	return user, nil
}

func (instance *repo) CountCompleteTasks(ctx context.Context, obj *ent.User) (int, error) {
	count, err := db.GetClient(ctx).Task.
		Query().
		Where(
			task.HasOwnerWith(user.ID(obj.ID)), // lấy task của user
			task.Completed(true),               // completed = true
		).
		Count(ctx)
	if err != nil {
		return 0, fmt.Errorf("userRepo.CompleteTasks: %w", err)
	}
	return count, nil
}

func (instance *repo) List(ctx context.Context) ([]*ent.User, error) {
	users, err := db.GetClient(ctx).User.Query().WithTasks().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("userRepo.List: %w", err)
	}

	return users, nil
}

func (instance *repo) ListUserTasks(ctx context.Context, obj *ent.User) ([]*ent.Task, error) {
	tasks, err := obj.QueryTasks().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("userRepo.ListTask: %d: %w", obj.ID, err)
	}
	return tasks, nil
}
