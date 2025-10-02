package user

import (
	"context"
	"fmt"
	"strconv"

	"go-demo/ent"
	"go-demo/internal/app/repository"
	"go-demo/internal/app/usecase/user/useriface"
	"go-demo/internal/pkg/factory/factoryiface"
	"go-demo/internal/pkg/graph/models"
)

type useCase struct {
	repo *repository.Repository
}

func NewUseCase(repoFactory factoryiface.RepositoryFactory) useriface.UseCase {
	return &useCase{
		repo: repoFactory.CreateRepository(),
	}
}

func (instance *useCase) Create(ctx context.Context, name string, email string) (*models.User, error) {
	user, err := instance.repo.User.Create(ctx, name, email)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return entToModel(user), nil
}

func (instance *useCase) List(ctx context.Context) ([]*models.User, error) {
	users, err := instance.repo.User.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}

	res := make([]*models.User, len(users))
	for i, user := range users {
		res[i] = entToModel(user)
	}

	return res, nil
}

func entToModel(entUser *ent.User) *models.User {
	model := &models.User{
		ID:    strconv.Itoa(entUser.ID),
		Name:  entUser.Name,
		Email: entUser.Email,
	}

	tasks := make([]*models.Task, len(entUser.Edges.Tasks))
	for i, task := range entUser.Edges.Tasks {
		tasks[i] = &models.Task{
			ID:        strconv.Itoa(task.ID),
			Title:     task.Title,
			Completed: task.Completed,
		}
	}

	model.Tasks = tasks

	return model
}
