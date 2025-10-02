package graph

import (
	"go-demo/internal/app/usecase/task/taskiface"
	"go-demo/internal/app/usecase/user/useriface"
	"go-demo/internal/pkg/factory/factoryiface"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserUseCase useriface.UseCase
	TaskUseCase taskiface.UseCase
}

func NewResolverRoot(useCaseFactory factoryiface.UseCaseFactory) *Resolver {
	return &Resolver{
		UserUseCase: useCaseFactory.CreateUserUseCase(),
		TaskUseCase: useCaseFactory.CreateTaskUseCase(),
	}
}
