package repository

import (
	"go-demo/internal/app/repository/task/taskiface"
	"go-demo/internal/app/repository/user/useriface"
)

type Repository struct {
	Task taskiface.Repository
	User useriface.Repository
}
