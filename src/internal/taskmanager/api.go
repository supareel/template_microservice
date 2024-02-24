package taskmanager

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskApplication interface {
	AddTask(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
	FindTaskById(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
	GetAllTasks(ctx *gin.Context)
}

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
//
//	@Summary		Add a task
//	@Description	used to add a task in db
//	@Tags			task
//	@Accept			json
//	@Produce		json
//
//	@Param			requestbody	body		AddTaskRequest	true	"Task object to be added"
//
//	@Success		201			{object}	TaskResponse
//	@Failure		400			{object}	TaskError
//	@Failure		404			{object}	TaskError
//	@Failure		500			{object}	TaskError
//	@Router			/api/v1/task [post]
func (s *TaskApi) AddTask(c *gin.Context) {
	var req AddTaskRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, TaskError{
			Message: err.Error(),
		})
		return
	}

	task, err := s.repo.CreateTask(c, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TaskError{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, TaskResponse{
		ID:        task.ID,
		Name:      task.Name,
		IsDone:    task.IsDone,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	})
}

// DeleteTask implements the DeleteTask method of ServerInterface
//
//	@Summary		Delete a task
//	@Description	delete task
//	@Tags			task
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of the task to delete"
//	@Success		200	{object}	DeleteTaskResponse
//	@Failure		400	{object}	TaskError
//	@Failure		404	{object}	TaskError
//	@Failure		500	{object}	TaskError
//	@Router			/api/v1/task/{id} [delete]
func (s *TaskApi) DeleteTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, TaskError{
			Message: "uuid is not in correct format",
		})
	}

	if err := s.repo.DeleteTaskById(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, TaskError{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusCreated, DeleteTaskResponse{
		IsDeleted: true,
	})
}

// FindTaskById implements the FindTaskById method of ServerInterface
//
//	@Summary		find task
//	@Description	get task by id
//	@Tags			task
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of the task to find"
//	@Success		200	{array}		TaskResponse
//	@Failure		400	{object}	TaskError
//	@Failure		404	{object}	TaskError
//	@Failure		500	{object}	TaskError
//	@Router			/api/v1/task/{id} [get]
func (s *TaskApi) FindTaskById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, TaskError{
			Message: "uuid is not in correct format",
		})
	}
	task, err := s.repo.FindTaskById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TaskError{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, TaskResponse{
		ID:        task.ID,
		Name:      task.Name,
		IsDone:    task.IsDone,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	})
}

// UpdateTask implements the UpdateTask method of ServerInterface
//
//	@Summary		update task
//	@Description	update task by id
//	@Tags			task
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path		string	true	"ID of the task to update"
//
// @Param			requestbody	body		UpdateTaskRequest	true	"Task object to be updated"
//
// @Success		200			{object}	TaskResponse
// @Failure		400			{object}	TaskError
// @Failure		404			{object}	TaskError
// @Failure		500			{object}	TaskError
// @Router			/api/v1/task/{id} [put]
func (s *TaskApi) UpdateTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, TaskError{
			Message: "uuid is not in correct format",
		})
	}
	var req UpdateTaskRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, TaskError{
			Message: err.Error(),
		})
		return
	}

	updatedtask, err := s.repo.UpdateTaskById(c, id, req.Name, req.IsDone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TaskError{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, TaskResponse{
		ID:        updatedtask.ID,
		Name:      updatedtask.Name,
		IsDone:    updatedtask.IsDone,
		CreatedAt: updatedtask.CreatedAt,
		UpdatedAt: updatedtask.UpdatedAt,
	})
}

// GetAllTasks implements the GetAllTasks method of ServerInterface
//
//	@Summary		List task
//	@Description	get task
//	@Tags			task
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		[]TaskResponse
//	@Failure		400	{object}	TaskError
//	@Failure		404	{object}	TaskError
//	@Failure		500	{object}	TaskError
//	@Router			/api/v1/task [get]
func (s *TaskApi) GetAllTasks(c *gin.Context) {
	allTask, err := s.repo.FindAllTask(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TaskError{
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, allTask)
}
