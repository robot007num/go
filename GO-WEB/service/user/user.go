package user

import (
	"github.com/robot007num/go/go-web/model/response"
	"github.com/robot007num/go/go-web/model/user"
	"github.com/robot007num/go/go-web/pkg/jwt"
	"github.com/robot007num/go/go-web/pkg/snowflake"
	sqluser "github.com/robot007num/go/go-web/repository/user"
)

const (
	logSuccess = "success"
	logError   = "error"
)

//RegisterService 用户注册逻辑
//string: 返回给客户端的状态码
//string: log日记显示失败或者成功
//string: log日记记录错误信息,返回""则代表无错误
func RegisterService(UserRe *user.Register) (response.ResCode, string, string) {
	//1. 验证之前是否已经注册
	ok, err := sqluser.VerifyUserExits(UserRe.Username)
	if err != nil {
		return response.CodeRegisterError, logError, response.InfoUserVerify
	}
	if ok {
		return response.CodeRegisterError, logError, response.InfoUserRegister
	}

	//2. 生成userid并存入数据库
	id, err := snowflake.CreateSnowID()
	if err != nil {
		return response.CodeRegisterError, logError, response.InfoUserSnowID
	}
	u := user.RegisterTable{
		Register: user.Register{
			Login: user.Login{
				Username: UserRe.Username,
				Password: UserRe.Password,
			},
		},
		Userid: id,
	}

	if err := sqluser.InsertUserTable(u); err != nil {
		return response.CodeRegisterError, logError, response.InfoUserInsert
	}

	return response.CodeRegisterSuccess, logSuccess, ""
}

func LoginService(UserRe *user.Login, allToken *jwt.AllToken) (response.ResCode, string, string) {
	//1. 验证账户是否存在
	ok, err := sqluser.VerifyUserExits(UserRe.Username)
	if err != nil {
		return response.CodeLoginError, logError, response.InfoUserVerify
	}
	if !ok {
		return response.CodeLoginError, logError, response.InfoUserUnRegister
	}

	//2. 存在此账户,则判断密码
	var sqlu user.Login
	if err := sqluser.VerifyUserLogin(UserRe.Username, &sqlu); err != nil {
		return response.CodeLoginError, logError, response.InfoUserSelect
	}

	if (UserRe.Username != sqlu.Username) || (UserRe.Password != sqlu.Password) {
		return response.CodeLoginError, logError, response.InfoUserPassword
	}

	//3. 生成access_token和refresh_token
	log := jwt.NewToken(sqlu, allToken)
	if log != "" {
		return response.CodeLoginError, logSuccess, log
	}

	return response.CodeLoginSuccess, logSuccess, ""
}
