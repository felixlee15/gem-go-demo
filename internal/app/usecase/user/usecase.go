package user

import (
	"context"
	"fmt"

	"gemdemo/ent"
	"gemdemo/internal/app/repository"
	"gemdemo/internal/app/usecase/user/useriface"
	"gemdemo/internal/pkg/factory/factoryiface"
)

type useCase struct {
	repo *repository.Repository
}

func NewUseCase(repoFactory factoryiface.RepositoryFactory) useriface.UseCase {
	return &useCase{
		repo: repoFactory.CreateRepository(),
	}
}

func (instance *useCase) Create(ctx context.Context, name string, email string) (*ent.User, error) {
	user, err := instance.repo.User.Create(ctx, name, email)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}

func (instance *useCase) CountCompleteTasks(ctx context.Context, obj *ent.User) (int, error) {
	count, err := instance.repo.User.CountCompleteTasks(ctx, obj)
	if err != nil {
		return 0, fmt.Errorf("count complete tasks: %w", err)
	}

	return count, nil
}

func (instance *useCase) List(ctx context.Context) ([]*ent.User, error) {
	users, err := instance.repo.User.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}

	return users, nil
}

func (instance *useCase) ListUserTasks(ctx context.Context, obj *ent.User) ([]*ent.Task, error) {
	tasks, err := instance.repo.User.ListUserTasks(ctx, obj)
	if err != nil {
		return nil, fmt.Errorf("list user tasks: %w", err)
	}

	return tasks, nil
}
