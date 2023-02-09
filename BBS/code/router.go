package code

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/bbs/router"
	"github.com/robot007num/go/bbs/utils"
	"net/http"
)

//初始化总路由

type server interface {
	ListenAndServe() error
}

func InitAllRouters() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System
	PublicGroup := Router.Group("")
	{
		//健康检测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
	}

	PrivateGroup := Router.Group("")
	PrivateGroup.Use(JwtAuth())
	{
		systemRouter.InitUserRouter(PrivateGroup)
	}
	return Router
}

const (
	ErrorTokenNil = "未登录/非法访问"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息

		//1. 检验是否有Token
		token := c.Request.Header.Get("x-token")
		if token == "" {
			utils.Result(-1, ErrorTokenNil, c)
			c.Abort()
			return
		}

		//2. 解析Token
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			utils.Result(-1, err.Error(), c)
			c.Abort()
			return
		}

		//3. 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", claims.UserName)
		c.Set("userid", claims.UserID)

		c.Next()
	}
}
