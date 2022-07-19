package taskmanager

import "time"

type Task struct {
	ID        string
	Name      string
	IsDone    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
