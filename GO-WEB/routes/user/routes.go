package user

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/go-web/controllers"
	"github.com/robot007num/go/go-web/utils"
)

var RegisterUserRoutes = func(e *gin.Engine) {

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)
	e.GET("/getsection", controllers.GetSection)
	e.GET("/getsection/:id", controllers.GetSectionClass)
	e.GET("/get-post/:classId", controllers.GetAllPost)
	e.GET("/get-post/:classId/:postId", controllers.GetSpecifyPost)

	v1 := e.Group("/v1").Use(utils.JWTAuthMiddleware())
	{
		v1.GET("/ping", controllers.Ping)
		v1.POST("/add-section", controllers.AddSection)
		v1.POST("/new-post", controllers.PostNewPost)

	}

	//e.POST("/changePassword", controllers.ChangePassword)
}
