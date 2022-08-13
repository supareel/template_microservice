package taskmanager

import (
	"context"
	"fmt"
	"gomicro/ent"
)

// business logic

// TaskService ...
type TaskService interface {
	GetAllTask(ctx context.Context) ([]*ent.Task, error)
	Create(ctx context.Context, name string) (Task, error)
}

// Task defines the application service in charge of interacting with Tasks.
type TaskSvc struct {
	repo TaskRepository
}

// NewTask ...
func NewTaskService(repo TaskRepository) *TaskSvc {
	return &TaskSvc{
		repo: repo,
	}
}

func (t *TaskSvc) GetAllTask(ctx context.Context) ([]*ent.Task, error) {
	tasks, err := t.repo.FindAll(ctx)
	if err != nil {
		return []*ent.Task{}, fmt.Errorf("repo create: %w", err)
	}

	return tasks, nil
}

// Create stores a new record.
func (t *TaskSvc) Create(ctx context.Context, name string) (Task, error) {
	// XXX: We will revisit the number of received arguments in future episodes.
	task, err := t.repo.Create(ctx, name)
	if err != nil {
		return Task{}, fmt.Errorf("repo create: %w", err)
	}

	return Task{
		ID:   task.ID,
		Name: task.Name,
	}, nil
}
