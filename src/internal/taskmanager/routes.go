package taskmanager

import "github.com/gin-gonic/gin"

func RegisterHandlers(router *gin.Engine, taskApi TaskApplication) {
	taskRouter := router.Group("/task")
	taskRouter.GET("/", taskApi.GetAllTasks)
	taskRouter.GET("/:id", taskApi.FindTaskById)
	taskRouter.POST("/:id", taskApi.AddTask)
	taskRouter.PUT("/:id", taskApi.UpdateTask)
	taskRouter.DELETE("/:id", taskApi.DeleteTask)
}
