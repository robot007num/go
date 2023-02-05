package code

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/router"
)

//初始化总路由

type server interface {
	ListenAndServe() error
}

func InitAllRouters() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System
	{
		systemRouter.InitUserRouter(Router)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
