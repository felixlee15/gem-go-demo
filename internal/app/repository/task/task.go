package task

import (
	"context"
	"strconv"
	"time"

	"go-demo/db"
	"go-demo/ent"
	"go-demo/ent/user"
	"go-demo/internal/app/repository/task/taskiface"
)

type repo struct{}

func NewRepository() taskiface.Repository {
	return &repo{}
}

func (instance *repo) Create(ctx context.Context, title string, userID string) (*ent.Task, error) {
	client := db.GetClient(ctx)

	uid, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}

	owner, err := client.User.Query().Where(user.ID(uid)).All(ctx)
	if err != nil {
		return nil, err
	}

	t, err := client.Task.Create().SetTitle(title).SetOwner(owner[0]).Save(ctx)
	return t, err
}

func (instance *repo) Update(ctx context.Context, taskID string, completed bool) (*ent.Task, error) {
	client := db.GetClient(ctx)

	taskIDInt, err := strconv.Atoi(taskID)
	if err != nil {
		return nil, err
	}

	task, err := client.Task.Get(ctx, taskIDInt)
	if err != nil {
		return nil, err
	}

	update := client.Task.UpdateOne(task).SetCompleted(completed)
	if completed {
		update.SetCompletedAt(time.Now())
	} else {
		update.ClearCompletedAt()
	}

	_, err = update.Save(ctx)
	if err != nil {
		return nil, err
	}

	owner, err := task.QueryOwner().All(ctx)
	if err != nil {
		return nil, err
	}

	task.Edges.Owner = owner[0]
	return task, err
}

func (instance *repo) List(ctx context.Context) ([]*ent.Task, error) {
	return db.GetClient(ctx).Task.Query().WithOwner().All(ctx)
}
