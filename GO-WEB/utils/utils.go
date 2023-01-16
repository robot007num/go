package utils

import (
	"strings"

	"github.com/robot007num/go/go-web/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/go-web/model/response"
	"github.com/robot007num/go/go-web/pkg/log"
)

//ParseBody 接收并检验参数
func ParseBody(c *gin.Context, x interface{}, info string) error {
	if err := c.ShouldBindJSON(x); err != nil {
		log.Info(info, log.String("result:", "error"),
			log.String("reason", "客户端传递参数错误"))
		return err
	}
	return nil
}

//JWTAuthMiddleware 验证Token有效性
//以后再根据具体情况补充日记
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			resStu := CreateReturnJson(response.CodeLoginError, response.InfoTokenNothing)
			ReturnBody(c, response.HttpOK, resStu)

			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			resStu := CreateReturnJson(response.CodeLoginError, response.InfoTokenFormat)
			ReturnBody(c, response.HttpOK, resStu)

			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParsingToken(parts[1])
		if err != nil {
			resStu := CreateReturnJson(response.CodeLoginError, response.InfoTokenInvalid)
			ReturnBody(c, response.HttpOK, resStu)

			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func ReturnBody(c *gin.Context, status int, res response.ReturnData) {
	c.JSON(status, res)
}

func RecordLog(program string, res string, info string) {
	log.Info(program, log.String("result:", res),
		log.String("reason", info))
}

func CreateReturnJson(code response.ResCode, data interface{}) response.ReturnData {
	res := response.ReturnData{
		Code: code,
		Msg:  code.Msg(),
		Data: data,
	}

	return res

}
