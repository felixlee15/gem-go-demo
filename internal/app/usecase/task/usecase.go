package task

import (
	"context"
	"fmt"

	"gemdemo/ent"
	"gemdemo/internal/app/repository"
	"gemdemo/internal/app/usecase/task/taskiface"
	"gemdemo/internal/pkg/factory/factoryiface"
)

type useCase struct {
	repo *repository.Repository
}

func NewUseCase(repoFactory factoryiface.RepositoryFactory) taskiface.UseCase {
	return &useCase{
		repo: repoFactory.CreateRepository(),
	}
}

func (instance *useCase) Create(ctx context.Context, title string, userID uint64) (*ent.Task, error) {
	t, err := instance.repo.Task.Create(ctx, title, userID)
	if err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}

	return t, nil
}

func (instance *useCase) Update(ctx context.Context, taskID uint64, completed bool) (*ent.Task, error) {
	t, err := instance.repo.Task.Update(ctx, taskID, completed)
	if err != nil {
		return nil, fmt.Errorf("update task: %w", err)
	}

	return t, nil
}

func (instance *useCase) List(ctx context.Context) ([]*ent.Task, error) {
	tasks, err := instance.repo.Task.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list tasks: %w", err)
	}

	return tasks, nil
}

func (instance *useCase) GetByID(ctx context.Context, id uint64) (*ent.Task, error) {
	task, err := instance.repo.Task.GetByID(ctx, id)
	if err != nil && ent.IsNotFound(err) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("get task: %w", err)
	}

	return task, nil
}

func (instance *useCase) GetByUser(ctx context.Context, userID uint64, completed *bool) ([]*ent.Task, error) {
	tasks, err := instance.repo.Task.GetByUser(ctx, userID, completed)
	if err != nil {
		return nil, fmt.Errorf("get tasks by user: %w", err)
	}

	return tasks, nil
}

func (instance *useCase) GetOwner(ctx context.Context, task *ent.Task) (*ent.User, error) {
	user, err := instance.repo.Task.GetOwner(ctx, task)
	if err != nil {
		return nil, fmt.Errorf("get owner of task: %w", err)
	}

	return user, nil
}
