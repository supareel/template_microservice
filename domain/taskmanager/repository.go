package taskmanager

import (
	"context"
	"fmt"
	"gomicro/database"

	"github.com/google/uuid"
)

// TaskRepository defines the datastore handling persisting Task records.
type TaskRepository interface {
	FindAll(ctx context.Context)([]Task, error)
	Create(ctx context.Context, name string) (Task, error)
	Find(ctx context.Context, id int32) (Task, error)
	Update(ctx context.Context, id int32, name string, isDone bool) error
}

type TaskRepo struct {
	db database.Database
}

func NewPostgresRepository(db database.Database) (*TaskRepo) {
	return &TaskRepo{
		db: db,
	}
}

func (trp *TaskRepo)	FindAll(ctx context.Context)([]Task, error) {
	client, err := trp.db.GetClient()
	if err != nil {
		fmt.Println(err)
		return []Task{}, err
	}
	resp, err := client.Task.Query().All(ctx)
	if err != nil {
		fmt.Println(err)
		return []Task{}, err
	}

	var taskList []Task
	for idx, task := range resp {
		taskList[idx] = Task{
			ID: task.ID.String(),
			Name: task.Name,
			IsDone: task.IsDone,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		}
	}
	return taskList, err
}

func (trp *TaskRepo) Create(ctx context.Context, name string) (Task, error){ 
	client, err := trp.db.GetClient()
	if err != nil {
		fmt.Println(err)
		return Task{}, err
	}
	resp, err := client.Task.Create().SetID(uuid.New()).SetName(name).SetIsDone(true).Save(ctx)
	if err != nil {
		fmt.Println(err)
		return Task{}, err
	}
	return Task{
		ID: resp.ID.String(),
		Name: resp.Name,
		IsDone: resp.IsDone,
		CreatedAt : resp.CreatedAt,
		UpdatedAt : resp.UpdatedAt,
	}, nil
}
func (t *TaskRepo) Find(ctx context.Context, id int32) (Task, error){ return Task{}, nil}
func (t *TaskRepo) Update(ctx context.Context, id int32, name string, isDone bool) error {return nil}