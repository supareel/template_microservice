package routes

import (
	"gomicro/controller"

	"github.com/gin-gonic/gin"
)

func LoginRoutes(mux *gin.Engine) {
	mux.POST("/login", controller.LoginUser)
}
