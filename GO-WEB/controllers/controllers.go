package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/go-web/model/response"
	"github.com/robot007num/go/go-web/model/user"
	"github.com/robot007num/go/go-web/pkg/jwt"
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
	var resStu response.ReturnData
	if err := utils.ParseBody(c, UserRe, GinRegister); err != nil {
		resStu = utils.CreateReturnJson(response.CodeInvalidParameters, err.Error())
		utils.ReturnBody(c, response.HttpOK, resStu)
		return
	}

	//2. 业务逻辑
	res, logres, info := service_user.RegisterService(UserRe)

	//3. 返回给客户端并记录此次结果
	resStu = utils.CreateReturnJson(res, info)
	utils.ReturnBody(c, response.HttpOK, resStu)

	//4. 记录日志
	utils.RecordLog(GinRegister, logres, res.Msg())

}

//LoginParameter 返回登录成功后生成的数据
type LoginParameter struct {
	Username     string `json:"username"`
	AccessToken  string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}

//Login 会返回两个Token 以后再优化下流程
func Login(c *gin.Context) {
	//1. 参数验证
	UserRe := &user.Login{}
	var resStu response.ReturnData
	if err := utils.ParseBody(c, UserRe, GinLogin); err != nil {
		utils.ReturnBody(c, response.HttpOK, resStu)
		return
	}

	//2. 业务逻辑
	Token := jwt.AllToken{}
	res, logres, info := service_user.LoginService(UserRe, &Token)

	//3. 返回给客户端并记录此次结果
	var l LoginParameter
	if info == "" {
		l.Username = UserRe.Username
		l.RefreshToken = Token.RefreshToken
		l.AccessToken = Token.AccessToken
		resStu = utils.CreateReturnJson(res, l)
	} else {
		resStu = utils.CreateReturnJson(res, info)
	}

	utils.ReturnBody(c, response.HttpOK, resStu)

	//4. 记录日志
	utils.RecordLog(GinLogin, logres, res.Msg())

}

func Ping(c *gin.Context) {
	name, _ := c.Get("username")
	RToken, _ := c.Get("RToken")
	AToken, _ := c.Get("AToken")
	fmt.Println(name.(string))
	l := LoginParameter{
		Username:     name.(string),
		RefreshToken: RToken.(string),
		AccessToken:  AToken.(string),
	}

	resStu := utils.CreateReturnJson(201, l)
	//3. 返回给客户端并记录此次结果
	utils.ReturnBody(c, response.HttpOK, resStu)
}
