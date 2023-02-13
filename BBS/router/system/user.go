package system

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/bbs/api/v1"
)

type UserRouter struct{}

func (*UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.POST("change_password", v1.ChangePassword)
	}
}
