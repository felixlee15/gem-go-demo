package useriface

import (
	"context"

	"go-demo/ent"
)

type Repository interface {
	Create(ctx context.Context, name string, email string) (*ent.User, error)
	List(ctx context.Context) ([]*ent.User, error)
}
