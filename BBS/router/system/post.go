package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/robot007num/go/bbs/api/v1"
)

type PostRouter struct{}

func (s *PostRouter) InitPostRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("post")
	{
		baseRouter.POST("new", v1.PostNew)
	}
	return baseRouter
}
