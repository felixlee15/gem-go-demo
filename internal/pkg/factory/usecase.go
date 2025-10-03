package factory

import (
	"gemdemo/internal/app/usecase/task"
	"gemdemo/internal/app/usecase/task/taskiface"
	"gemdemo/internal/app/usecase/user"
	"gemdemo/internal/app/usecase/user/useriface"
	"gemdemo/internal/pkg/factory/factoryiface"
)

type useCaseFactory struct {
	repoFactory factoryiface.RepositoryFactory
	user        useriface.UseCase
	task        taskiface.UseCase
}

func (u *useCaseFactory) CreateUserUseCase() useriface.UseCase {
	if u.user == nil {
		u.user = user.NewUseCase(u.repoFactory)
	}

	return u.user
}

func (u *useCaseFactory) CreateTaskUseCase() taskiface.UseCase {
	if u.task == nil {
		u.task = task.NewUseCase(u.repoFactory)
	}

	return u.task
}

func NewUseCaseFactory() factoryiface.UseCaseFactory {
	return &useCaseFactory{
		repoFactory: NewRepositoryFactory(),
	}
}
