package taskmanager

import (
	"context"
	"fmt"
)

// business logic

// TaskService ...
type TaskService interface {
	Create(ctx context.Context, name string) (Task, error)
	Task(ctx context.Context, id string) (Task, error)
	Update(ctx context.Context, id string, name string, isDone bool) error
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



// Create stores a new record.
func (t *TaskSvc) Create(ctx context.Context, name string) (Task, error) {
	// XXX: We will revisit the number of received arguments in future episodes.
	task, err := t.repo.Create(ctx, name)
	if err != nil {
		return Task{}, fmt.Errorf("repo create: %w", err)
	}

	return Task{
		ID: task.ID,
		Name: task.Name,
	}, nil
}

// Task gets an existing Task from the datastore.
func (t *TaskSvc) Task(ctx context.Context, id string) (Task, error) {
	// XXX: We will revisit the number of received arguments in future episodes.
	task, err := t.repo.Find(ctx, id)
	if err != nil {
		return Task{}, fmt.Errorf("repo find: %w", err)
	}

	return Task{
		ID: task.ID,
		Name: task.Name,
	}, nil
}

// Update updates an existing Task in the datastore.
func (t *TaskSvc) Update(ctx context.Context, id string, name string, isDone bool) error {
	// XXX: We will revisit the number of received arguments in future episodes.
	if err := t.repo.Update(ctx, id, name, isDone); err != nil {
		return fmt.Errorf("repo update: %w", err)
	}

	return nil
}