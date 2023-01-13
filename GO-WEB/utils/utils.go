package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/go-web/pkg/log"
	"net/http"
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

func ReturnBody(c *gin.Context, res string) {
	c.JSON(http.StatusOK, gin.H{
		"msg": res,
	})
}

func RecordLog(program string, res string, info string) {
	log.Info(program, log.String("result:", res),
		log.String("reason", info))
}
