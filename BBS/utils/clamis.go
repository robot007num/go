package utils

import (
	"github.com/gin-gonic/gin"
)

// GetUserName 从Gin的Context中获取从jwt解析出来的用户名
func GetUserName(c *gin.Context) string {
	username, _ := c.Get("username")
	return username.(string)
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserUuid(c *gin.Context) int64 {
	userid, _ := c.Get("userid")
	return userid.(int64)
}
