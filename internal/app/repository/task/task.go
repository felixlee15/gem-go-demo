package task

import (
	"context"
	"fmt"
	"time"

	"gemdemo/db"
	"gemdemo/ent"
	"gemdemo/ent/task"
	"gemdemo/ent/user"
	"gemdemo/internal/app/repository/task/taskiface"
)

type repo struct{}

func NewRepository() taskiface.Repository {
	return &repo{}
}

func (instance *repo) Create(ctx context.Context, title string, userID uint64) (*ent.Task, error) {
	client := db.GetClient(ctx)

	owner, err := client.User.Query().Where(user.ID(userID)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	if owner == nil {
		return nil, fmt.Errorf("owner not found")
	}

	t, err := client.Task.Create().SetTitle(title).SetOwner(owner).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("taskRepo.Create: %w", err)
	}

	return t, nil
}

func (instance *repo) Update(ctx context.Context, taskID uint64, completed bool) (*ent.Task, error) {
	client := db.GetClient(ctx)

	t, err := client.Task.Get(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("get task: %w", err)
	}

	update := client.Task.UpdateOne(t).SetCompleted(completed)
	if completed {
		update.SetCompletedAt(time.Now())
	} else {
		update.ClearCompletedAt()
	}

	_, err = update.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("taskRepo.Update: %w", err)
	}

	owner, err := t.QueryOwner().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("get owner: %w", err)
	}

	t.Edges.Owner = owner[0]
	return t, err
}

func (instance *repo) List(ctx context.Context) ([]*ent.Task, error) {
	tasks, err := db.GetClient(ctx).Task.Query().WithOwner().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("taskRepo.List: %w", err)
	}

	return tasks, nil
}

func (instance *repo) GetByID(ctx context.Context, id uint64) (*ent.Task, error) {
	t, err := db.GetClient(ctx).Task.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("taskRepo.GetByID %d: %w", id, err)
	}

	return t, nil
}

func (instance *repo) GetByUser(ctx context.Context, userID uint64, completed *bool) ([]*ent.Task, error) {
	query := db.GetClient(ctx).Task.Query().Where(task.OwnerID(userID))

	if completed != nil {
		query = query.Where(task.Completed(*completed))
	}

	tasks, err := query.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("taskRepo.GetByUser: %d: %w", userID, err)
	}

	return tasks, nil
}

func (instance *repo) GetOwner(ctx context.Context, obj *ent.Task) (*ent.User, error) {
	owner, err := obj.QueryOwner().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("taskRepo.GetOwner: %d: %w", obj.ID, err)
	}

	return owner, nil
}
