package taskmanager

import (
	"time"

	"github.com/google/uuid"
)

type AddTaskRequest struct {
	Name string `json:"name" binding:"required"`
}

type AddTaskResponse struct {
	Name   string `json:"name"`
	IsDone bool   `json:"isDone"`
}

type UpdateTaskRequest struct {
	Name   *string `json:"name,omitempty" binding:"required"`
	IsDone *bool   `json:"isDone,omitempty" binding:"required"`
}

type DeleteTaskResponse struct {
	IsDeleted bool `json:"isDeleted"`
}

type TaskResponse struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IsDone holds the value of the "is_done" field.
	IsDone bool `json:"is_done,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
