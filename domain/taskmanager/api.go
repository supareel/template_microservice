package taskmanager

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TaskHandler ...
type TaskHandler struct {
	svc TaskService
}

// NewTaskHandler ...
func NewTaskHandler(svc TaskService) *TaskHandler {
	return &TaskHandler{
		svc: svc,
	}
}


func (t *TaskHandler) create(c *gin.Context) {
	var req CreateTasksRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
			"hint":  "please check if task name is valid",
		})
		return
	}

	task, err := t.svc.Create(c.Request.Context(), req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "create failed",
			"hint":  "please check if task name is valid",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "",
		"id":      task.ID,
		"name":    task.Name,
	})
	return
}

func (t *TaskHandler) task(c *gin.Context) {

	id, _ := c.Params.Get("id") // NOTE: Safe to ignore error, because it's always defined.

	task, err := t.svc.Task(c.Request.Context(), id)
	if err != nil {
		// XXX: Differentiating between NotFound and Internal errors will be covered in future episodes.
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "find failed",
			"hint":  "please check if task name is valid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"id":      task.ID,
		"name":    task.Name,
	})
}

// UpdateTasksRequest defines the request used for updating a task.

func (t *TaskHandler) update(c *gin.Context) {
	var req UpdateTasksRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
			"hint":  "please check if task name is valid",
		})
		return
	}

	id, _ := c.Params.Get("id") // NOTE: Safe to ignore error, because it's always defined.

	err := t.svc.Update(c.Request.Context(), id, req.Name, req.IsDone)
	if err != nil {
		// XXX: Differentiating between NotFound and Internal errors will be covered in future episodes.
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "update failed",
			"hint":  "please check if task name is valid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update successful",
	})
}
