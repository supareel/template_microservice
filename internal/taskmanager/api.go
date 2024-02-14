package taskmanager

import (
	taskmanager_gen "gomicro/internal/taskmanager/_generated_"

	"github.com/gin-gonic/gin"
)

// Task defines the application service in charge of interacting with Tasks.
type TaskApi struct {
	repo TaskRepository
}

// NewTask ...
func NewTaskApi(repo TaskRepository) *TaskApi {
	return &TaskApi{
		repo: repo,
	}
}

// AddTask implements the AddTask method of ServerInterface
func (s *TaskApi) AddTask(c *gin.Context) {
	// Your implementation for adding a task
}

// DeleteTask implements the DeleteTask method of ServerInterface
func (s *TaskApi) DeleteTask(c *gin.Context, id int64) {
	// Your implementation for deleting a task by ID
}

// FindTaskById implements the FindTaskById method of ServerInterface
func (s *TaskApi) FindTaskById(c *gin.Context, id int64) {
	// Your implementation for finding a task by ID
}

// UpdateTask implements the UpdateTask method of ServerInterface
func (s *TaskApi) UpdateTask(c *gin.Context, id int64) {
	// Your implementation for updating a task by ID
}

// GetAllTasks implements the GetAllTasks method of ServerInterface
func (s *TaskApi) GetAllTasks(c *gin.Context, params taskmanager_gen.GetAllTasksParams) {
	// Your implementation for getting all tasks
}
