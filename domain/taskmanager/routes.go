package taskmanager

import (
	"github.com/gin-gonic/gin"
)

const uuidRegEx string = `[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}`

// Register connects the handlers to the router.
func (t *TaskHandler) Register(r *gin.Engine) {
	r.POST("/tasks", t.create)
	r.GET("/tasks/:id", t.task)
	r.PUT("/tasks/:id", t.update)
}
