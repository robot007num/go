package system

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/bbs/api/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", v1.Login)
		baseRouter.POST("register", v1.Register)
		baseRouter.GET("refresh_token", v1.RefreshToken) //只用于刷新AccessToken
		baseRouter.GET("allCommunity", v1.GetAllCommunity)
		baseRouter.GET("section/:communityID", v1.GetSection)
		baseRouter.GET("post/:postId", v1.PostGet)
		baseRouter.GET("section/:communityID/post", v1.GetSectionPost)
	}
	return baseRouter
}
