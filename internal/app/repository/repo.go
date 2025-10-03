package repository

import (
	"gemdemo/internal/app/repository/task/taskiface"
	"gemdemo/internal/app/repository/user/useriface"
)

type Repository struct {
	Task taskiface.Repository
	User useriface.Repository
}
