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
	FindAllTask(ctx context.Context) ([]*ent.Task, error)
	FindTaskById(ctx context.Context, id uuid.UUID) (*ent.Task, error)
	CreateTask(ctx context.Context, name string) (*ent.Task, error)
	UpdateTaskById(ctx context.Context, id uuid.UUID, name *string, isDone *bool) (*ent.Task, error)
	DeleteTaskById(ctx context.Context, id uuid.UUID) error
}

type TaskRepo struct {
	db database.Database
}

func NewPostgresRepository(db database.Database) *TaskRepo {
	return &TaskRepo{
		db: db,
	}
}

func (trp *TaskRepo) FindAllTask(ctx context.Context) ([]*ent.Task, error) {
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

func (trp *TaskRepo) FindTaskById(ctx context.Context, id uuid.UUID) (*ent.Task, error) {
	client, err := trp.db.GetClient()
	if err != nil {
		fmt.Println(err)
		return &ent.Task{}, err
	}
	resp, err := client.Task.Get(ctx, id)
	if err != nil {
		pkg.FancyHandleError(err)
		return &ent.Task{}, err
	}
	return resp, err
}

func (trp *TaskRepo) CreateTask(ctx context.Context, name string) (*ent.Task, error) {
	client, err := trp.db.GetClient()
	if err != nil {
		fmt.Println(err)
		return &ent.Task{}, err
	}
	resp, err := client.Task.Create().
		SetID(uuid.New()).
		SetName(name).
		SetIsDone(false).
		Save(ctx)
	if err != nil {
		fmt.Println(err)
		return &ent.Task{}, err
	}

	return resp, nil
}

func (trp *TaskRepo) UpdateTaskById(ctx context.Context, id uuid.UUID, name *string, isDone *bool) (*ent.Task, error) {
	client, err := trp.db.GetClient()
	if err != nil {
		fmt.Println(err)
		return &ent.Task{}, err
	}

	query := client.Task.UpdateOneID(id)

	if name != nil {
		query = query.SetName(*name)
	}
	if isDone != nil {
		query = query.SetIsDone(true)
	}

	resp, err := query.Save(ctx)
	if err != nil {
		fmt.Println(err)
		return &ent.Task{}, err
	}

	return resp, nil
}

func (trp *TaskRepo) DeleteTaskById(ctx context.Context, id uuid.UUID) error {
	client, err := trp.db.GetClient()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = client.Task.DeleteOneID(id).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
