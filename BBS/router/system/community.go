package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/robot007num/go/bbs/api/v1"
)

type CommunityRouter struct{}

func (s *CommunityRouter) InitCommunityRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("community")
	{
		baseRouter.POST("add", v1.AddCommunity)
	}
	return baseRouter
}
