package factory

import (
	"go-demo/internal/app/repository"
	"go-demo/internal/app/repository/task"
	"go-demo/internal/app/repository/user"
)

type RepositoryFactory struct{}

func (RepositoryFactory) CreateRepository() *repository.Repository {
	return &repository.Repository{
		User: user.NewRepository(),
		Task: task.NewRepository(),
	}
}

func NewRepositoryFactory() *RepositoryFactory {
	return &RepositoryFactory{}
}
