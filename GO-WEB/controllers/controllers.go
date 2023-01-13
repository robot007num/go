package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/go-web/model/user"
	service_user "github.com/robot007num/go/go-web/service/user"
	"github.com/robot007num/go/go-web/utils"
)

const (
	GinRegister = "[gin/Register]"
	GinLogin    = "[gin/login]"
)

func Register(c *gin.Context) {
	//1. 参数验证
	UserRe := &user.Register{}
	if err := utils.ParseBody(c, UserRe, "[gin/Register]"); err != nil {
		utils.ReturnBody(c, err.Error())
		return
	}

	//2. 业务逻辑
	res, logres, info := service_user.RegisterService(UserRe)

	//3. 返回给客户端并记录此次结果
	utils.ReturnBody(c, res)

	//4. 记录日志
	utils.RecordLog(GinRegister, logres, info)

}

func Login(c *gin.Context) {
	//1. 参数验证
	UserRe := &user.Login{}
	if err := utils.ParseBody(c, UserRe, "[gin/Login]"); err != nil {
		utils.ReturnBody(c, err.Error())
		return
	}

	//2. 业务逻辑
	res, logres, info := service_user.LoginService(UserRe)

	//3. 返回给客户端并记录此次结果
	utils.ReturnBody(c, res)

	//4. 记录日志
	utils.RecordLog(GinLogin, logres, info)

}

func ChangePassword(c *gin.Context) {

}
