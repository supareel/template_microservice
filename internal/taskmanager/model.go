package taskmanager

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID
	Name      string
	IsDone    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
