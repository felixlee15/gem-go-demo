package factory

import (
	"gemdemo/internal/app/repository"
	"gemdemo/internal/app/repository/task"
	"gemdemo/internal/app/repository/user"
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
