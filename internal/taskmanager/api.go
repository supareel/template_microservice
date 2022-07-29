package taskmanager

import (
	"context"
	proto "gomicro/pkg/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// TaskHandler ...
type TaskHandler struct {
	svc TaskService
	proto.UnimplementedTaskServiceServer
}

// NewTaskHandler ...
func NewTaskHandler(svc TaskService) *TaskHandler {
	return &TaskHandler{
		svc: svc,
	}
}

func (t *TaskHandler) GetAllTasks(ctx context.Context, in *proto.GetAllTasksRequest) (*proto.GetAllTasksResponse, error) {
	tasks, err := t.svc.GetAllTask(ctx)
	if err != nil {
		return nil, err
	}
	newTasksResponse := proto.GetAllTasksResponse{
		Data:   make([]*proto.Task, len(tasks)),
		Status: proto.DBOPERATIONSTATUS_QUERIED,
	}


	for idx, task := range tasks {
		newTasksResponse.Data[idx] = &proto.Task{
			Name:   task.Name,
			Isdone: task.IsDone,
		}
	}

	return &newTasksResponse, nil
}

func (t *TaskHandler) CreateTask(ctx context.Context, in *proto.CreateTasksRequest) (*proto.CreateTasksResponse, error) {
	task, err := t.svc.Create(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	taskResponse := &proto.CreateTasksResponse{
		Data: &proto.Task{
			Name:      task.Name,
			Isdone:    task.IsDone,
			Createdat: timestamppb.New(task.CreatedAt),
			Updatedat: timestamppb.New(task.UpdatedAt),
		},
		Status: proto.DBOPERATIONSTATUS_CREATED,
	}
	return taskResponse, nil
}