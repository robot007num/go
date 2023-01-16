package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/go-web/model/response"
	"github.com/robot007num/go/go-web/model/user"
	service_user "github.com/robot007num/go/go-web/service/user"
	"github.com/robot007num/go/go-web/utils"
	"net/http"
)

const (
	GinRegister = "[gin/Register]"
	GinLogin    = "[gin/login]"
)

var (
	Sok = http.StatusOK
)

func createReturnJson(code response.ResCode, data interface{}) response.ReturnData {
	res := response.ReturnData{
		Code: code,
		Msg:  code.Msg(),
		Data: data,
	}

	return res

}

func Register(c *gin.Context) {
	//1. 参数验证
	UserRe := &user.Register{}
	var resStu response.ReturnData
	if err := utils.ParseBody(c, UserRe, GinRegister); err != nil {
		resStu = createReturnJson(response.CodeInvalidParameters, err.Error())
		utils.ReturnBody(c, Sok, resStu)
		return
	}

	//2. 业务逻辑
	res, logres, info := service_user.RegisterService(UserRe)

	resStu = createReturnJson(res, info)
	//3. 返回给客户端并记录此次结果
	utils.ReturnBody(c, Sok, resStu)

	//4. 记录日志
	utils.RecordLog(GinRegister, logres, res.Msg())

}

func Login(c *gin.Context) {
	//1. 参数验证
	UserRe := &user.Login{}
	var resStu response.ReturnData
	if err := utils.ParseBody(c, UserRe, GinLogin); err != nil {
		utils.ReturnBody(c, Sok, resStu)
		return
	}

	//2. 业务逻辑
	res, logres, info := service_user.LoginService(UserRe)

	resStu = createReturnJson(res, info)
	//3. 返回给客户端并记录此次结果
	utils.ReturnBody(c, Sok, resStu)

	//4. 记录日志
	utils.RecordLog(GinLogin, logres, res.Msg())

}
