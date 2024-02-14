package taskmanager

import (
	"context"
	"fmt"
	"gomicro/src/database"
	"gomicro/src/ent"
	"gomicro/src/pkg"

	"github.com/google/uuid"
)

// TaskRepository defines the datastore handling persisting Task records.
type TaskRepository interface {
	FindAll(ctx context.Context) ([]*ent.Task, error)
	Create(ctx context.Context, name string) (ent.Task, error)
}

type TaskRepo struct {
	db database.Database
}

func NewPostgresRepository(db database.Database) *TaskRepo {
	return &TaskRepo{
		db: db,
	}
}

func (trp *TaskRepo) FindAll(ctx context.Context) ([]*ent.Task, error) {
	client, err := trp.db.GetClient()
	if err != nil {
		fmt.Println(err)
		return []*ent.Task{}, err
	}
	resp, err := client.Task.Query().All(ctx)
	if err != nil {
		pkg.FancyHandleError(err)
		return []*ent.Task{}, err
	}

	return resp, err
}

func (trp *TaskRepo) Create(ctx context.Context, name string) (ent.Task, error) {
	client, err := trp.db.GetClient()
	if err != nil {
		fmt.Println(err)
		return ent.Task{}, err
	}
	resp, err := client.Task.Create().SetID(uuid.New()).SetName(name).SetIsDone(true).Save(ctx)
	if err != nil {
		fmt.Println(err)
		return ent.Task{}, err
	}

	return *resp, nil
}
