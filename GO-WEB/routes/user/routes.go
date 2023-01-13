package user

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/go-web/controllers"
)

var RegisterUserRoutes = func(e *gin.Engine) {
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)
	e.POST("/changePassword", controllers.ChangePassword)
}
