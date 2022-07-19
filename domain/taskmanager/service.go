package taskmanager

import (
	"context"
	"fmt"
)

// business logic

// TaskService ...
type TaskService interface {
	GetAllTask(ctx context.Context) ([]Task, error)
	Create(ctx context.Context, name string) (Task, error)
	Task(ctx context.Context, id int32) (Task, error)
	Update(ctx context.Context, id int32, name string, isDone bool) error
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

func (t *TaskSvc) GetAllTask(ctx context.Context) ([]Task, error) {
	tasks, err := t.repo.FindAll(ctx)
	if err != nil {
		return []Task{}, fmt.Errorf("repo create: %w", err)
	}

	var allTask []Task

	for idx, task := range tasks {
		allTask[idx] = Task{
			ID:   task.ID,
			Name: task.Name,
		}
	}

	return allTask, nil
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

// Task gets an existing Task from the datastore.
func (t *TaskSvc) Task(ctx context.Context, id int32) (Task, error) {
	// XXX: We will revisit the number of received arguments in future episodes.
	task, err := t.repo.Find(ctx, id)
	if err != nil {
		return Task{}, fmt.Errorf("repo find: %w", err)
	}

	return Task{
		ID:   task.ID,
		Name: task.Name,
	}, nil
}

// Update updates an existing Task in the datastore.
func (t *TaskSvc) Update(ctx context.Context, id int32, name string, isDone bool) error {
	// XXX: We will revisit the number of received arguments in future episodes.
	if err := t.repo.Update(ctx, id, name, isDone); err != nil {
		return fmt.Errorf("repo update: %w", err)
	}

	return nil
}
