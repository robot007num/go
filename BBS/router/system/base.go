package system

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/bbs/controller"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", controller.Login)
		baseRouter.POST("register", controller.Register)
	}
	return baseRouter
}
