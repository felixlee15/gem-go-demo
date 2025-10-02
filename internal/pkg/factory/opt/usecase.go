package factory

import (
	"go-demo/internal/app/usecase/task"
	"go-demo/internal/app/usecase/task/taskiface"
	"go-demo/internal/app/usecase/user"
	"go-demo/internal/app/usecase/user/useriface"
	"go-demo/internal/pkg/factory/factoryiface"
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
