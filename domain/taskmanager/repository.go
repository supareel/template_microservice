package taskmanager

import (
	"context"
	"fmt"
	"gomicro/database"

	"github.com/google/uuid"
)

// TaskRepository defines the datastore handling persisting Task records.
type TaskRepository interface {
	Create(ctx context.Context, name string) (TaskModel, error)
	Find(ctx context.Context, id string) (TaskModel, error)
	Update(ctx context.Context, id string, name string, isDone bool) error
}

type TaskRepo struct {
	db database.Database
}

func NewPostgresRepository(db database.Database) (*TaskRepo) {
	return &TaskRepo{
		db: db,
	}
}


func (trp *TaskRepo) Create(ctx context.Context, name string) (TaskModel, error){ 
	client, err := trp.db.GetClient()
	resp, err := client.Task.Create().SetID(uuid.New()).SetName(name).SetIsDone(true).Save(ctx)
	if err != nil {
		fmt.Println(err)
		return TaskModel{}, err
	}
	return TaskModel{
		ID: resp.ID.String(),
		Name: resp.Name,
		IsDone: resp.IsDone,
		CreatedAt : resp.CreatedAt,
		UpdatedAt : resp.UpdatedAt,
	}, nil
}
func (t *TaskRepo) Find(ctx context.Context, id string) (TaskModel, error){ return TaskModel{}, nil}
func (t *TaskRepo) Update(ctx context.Context, id string, name string, isDone bool) error {return nil}