package factoryiface

import (
	"go-demo/internal/app/repository"
	"go-demo/internal/app/usecase/task/taskiface"
	"go-demo/internal/app/usecase/user/useriface"
)

type UseCaseFactory interface {
	CreateUserUseCase() useriface.UseCase
	CreateTaskUseCase() taskiface.UseCase
}

type RepositoryFactory interface {
	CreateRepository() *repository.Repository
}
