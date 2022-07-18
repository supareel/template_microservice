package taskmanager

import "time"

type TaskModel struct {
	ID        string
	Name      string
	IsDone    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Task is an activity that needs to be completed within a period of time.
type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateTasksRequest defines the request used for creating tasks.
type CreateTasksRequest struct {
	Name string `json:"name"`
	// XXX `Priority` and `Dates` are intentionally missing, to be covered in future videos
}

// CreateTasksResponse defines the response returned back after creating tasks.
type CreateTasksResponse struct {
	Task Task `json:"task"`
}

type UpdateTasksRequest struct {
	Name string `json:"description"`
	IsDone      bool   `json:"is_done"`
	// XXX `Priority` and `Dates` are intentionally missing, to be covered in future videos
}

// GetTasksResponse defines the response returned back after searching one task.
type GetTasksResponse struct {
	Task Task `json:"task"`
}