package code

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/bbs/global"
	"go.uber.org/zap"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func RunWindowsServer() {
	//初始化Redis

	//
	Router := InitAllRouters()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	global.GVA_LOG.Info(s.ListenAndServe().Error())
}
