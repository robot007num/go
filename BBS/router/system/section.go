package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/robot007num/go/bbs/api/v1"
)

type SectionRouter struct{}

func (s *SectionRouter) InitSectionRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("section")
	{
		baseRouter.POST("add", v1.AddSection)
	}
	return baseRouter
}
