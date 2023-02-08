package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/bbs/model/common/response"
	"net/http"
	"os"
)

//@author: [robot007num]
//@function: PathExists
//@description: 文件目录是否存在
//@param: path string
//@return: bool, error

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Result(code response.ResCode, data interface{}, c *gin.Context) {
	res := response.Response{
		Code: int64(code),
		Msg:  code.Msg(),
		Data: data,
	}

	c.JSON(http.StatusOK, res)
}
