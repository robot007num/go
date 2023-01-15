package main

import (
	"github.com/robot007num/go/go-web/pkg/gin"
	"github.com/robot007num/go/go-web/pkg/log"
	"github.com/robot007num/go/go-web/pkg/viper"
	"github.com/robot007num/go/go-web/repository"
)

func main() {
	//1. 读取配置文件
	viper.InitViper()

	//2. 开启数据库
	repository.StartMySql()
	if err := repository.SelectTableIsExits("user"); err != nil {
		repository.CreateTableFromSQLFile()
	}

	//3. 开启Redis
	repository.StartRedis()

	//4. 初始化日志
	log.InitLog()
	defer log.Sync()

	//5. 记录日记
	log.Info("[Program]", log.String("result", "success"),
		log.String("reason", "环境初始化已全部完成"))
	log.Info("[Program]", log.String("result", " Program start"))

	//6. 开启gin服务
	gin.InitGin()

}
