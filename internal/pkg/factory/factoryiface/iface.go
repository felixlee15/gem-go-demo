package factoryiface

import (
	"gemdemo/internal/app/repository"
	"gemdemo/internal/app/usecase/task/taskiface"
	"gemdemo/internal/app/usecase/user/useriface"
)

type UseCaseFactory interface {
	CreateUserUseCase() useriface.UseCase
	CreateTaskUseCase() taskiface.UseCase
}

type RepositoryFactory interface {
	CreateRepository() *repository.Repository
}
