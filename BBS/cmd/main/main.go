package main

import (
	"github.com/robot007num/go/bbs/code"
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/pkg"
	"go.uber.org/zap"
	"os"
)

func main() {

	//读取配置文件
	global.GVA_VIPER = code.Viper()
	//初始化Zap日志库
	global.GVA_LOG = code.Zap()
	defer global.GVA_LOG.Sync()

	//初始化mysql数据库
	global.GVA_DB = pkg.Sqlx()
	if global.GVA_DB == nil {
		//global.GVA_LOG.Info("服务器退出", zap.String("初始化数据库", "失败"))
		global.GVA_LOG.Error("服务器退出", zap.String("初始化数据库", "失败"))
		os.Exit(-1)
	}

	//开启web服务
	code.RunWindowsServer()
}

/*
2023年2月3日 23点46分
因为对Zap库封装不熟 导致了我既想要大白的多个文件写入也想要按照日期创建目录并写入日志。
但现在只能满足一个 所以还待了解。
然后今天发现一本《程序员思维修炼》 打算看完后再搞Zap库
*/
