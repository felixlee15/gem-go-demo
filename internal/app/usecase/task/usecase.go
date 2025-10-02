package task

import (
	"context"
	"fmt"
	"strconv"

	"go-demo/ent"
	"go-demo/internal/app/repository"
	"go-demo/internal/app/usecase/task/taskiface"
	"go-demo/internal/pkg/factory/factoryiface"
	"go-demo/internal/pkg/graph/models"
)

type useCase struct {
	repo *repository.Repository
}

func NewUseCase(repoFactory factoryiface.RepositoryFactory) taskiface.UseCase {
	return &useCase{
		repo: repoFactory.CreateRepository(),
	}
}

func (instance *useCase) Create(ctx context.Context, title string, userID string) (*models.Task, error) {
	t, err := instance.repo.Task.Create(ctx, title, userID)
	if err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}

	return entToModel(t), nil
}

func (instance *useCase) Update(ctx context.Context, taskID string, completed bool) (*models.Task, error) {
	t, err := instance.repo.Task.Update(ctx, taskID, completed)
	if err != nil {
		return nil, fmt.Errorf("update task: %w", err)
	}

	return entToModel(t), nil
}

func (instance *useCase) List(ctx context.Context) ([]*models.Task, error) {
	tasks, err := instance.repo.Task.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list tasks: %w", err)
	}

	res := make([]*models.Task, len(tasks))
	for i, task := range tasks {
		res[i] = entToModel(task)
	}

	return res, nil
}

func entToModel(entTask *ent.Task) *models.Task {
	model := &models.Task{
		ID:        strconv.Itoa(entTask.ID),
		Title:     entTask.Title,
		Completed: entTask.Completed,
	}

	if entTask.Edges.Owner != nil {
		model.Owner = &models.User{
			ID:    strconv.Itoa(entTask.Edges.Owner.ID),
			Name:  entTask.Edges.Owner.Name,
			Email: entTask.Edges.Owner.Email,
		}
	}

	return model
}
