package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/go-web/model/config"
	"github.com/robot007num/go/go-web/routes/user"
	"io"
	"os"
)

//封装ZAP库，再使用Gin中间件接收Gin日记没用(?)

func InitGin() {

	f, _ := os.Create(config.GetAllConfig().Log.GinFileName)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	user.RegisterUserRoutes(r)

	//r.Use(GinLogger(log.Default()), GinRecovery(log.Default(), true))
	r.Run(":8080")

}
