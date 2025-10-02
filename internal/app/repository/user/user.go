package user

import (
	"context"

	"go-demo/db"
	"go-demo/ent"
	"go-demo/internal/app/repository/user/useriface"
)

type repo struct{}

func NewRepository() useriface.Repository {
	return &repo{}
}

func (instance *repo) Create(ctx context.Context, name string, email string) (*ent.User, error) {
	return db.GetClient(ctx).User.Create().SetName(name).SetEmail(email).Save(ctx)
}

func (instance *repo) List(ctx context.Context) ([]*ent.User, error) {
	return db.GetClient(ctx).User.Query().WithTasks().All(ctx)
}
