package user

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/go-web/controllers"
	"github.com/robot007num/go/go-web/utils"
)

var RegisterUserRoutes = func(e *gin.Engine) {
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)
	e.GET("/ping", utils.JWTAuthMiddleware(), controllers.Ping)
	e.POST("/addsection", controllers.AddSection)
	e.GET("/getsection", controllers.GetSection)
	e.GET("/getsection/:id", controllers.GetSectionClass)
	//e.POST("/changePassword", controllers.ChangePassword)
}
