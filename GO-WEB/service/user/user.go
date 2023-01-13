package user

import (
	"github.com/robot007num/go/go-web/model/user"
	"github.com/robot007num/go/go-web/pkg/snowflake"
	sqluser "github.com/robot007num/go/go-web/repository/user"
)

const (
	errRegister = "该用户已注册"
	errVerify   = "验证用户失败"
	errSnowID   = "生成用户ID失败"
	errInsert   = "插入新用户失败"
	errLogin    = "该用户未注册"
	errSelect   = "查询用户失败"
	errPassword = "密码错误"

	logSuccess = "success"
	logError   = "error"
)

//返回给客户端
const (
	returnClientRegisterErr     = "注册失败"
	returnClientRegisterSuccess = "注册成功"
	returnClientLoginErr        = "登录失败"
	returnClientLoginSuccess    = "登录成功"
)

//RegisterService 用户注册逻辑
//string: 返回给客户端的信息
//string: 日记需求
//string: 记录日记
func RegisterService(UserRe *user.Register) (string, string, string) {
	//1. 验证之前是否已经注册
	ok, err := sqluser.VerifyUserExits(UserRe.Username)
	if err != nil {
		return returnClientRegisterErr, logError, errVerify
	}
	if ok {
		return errRegister, logError, errRegister
	}

	//2. 生成userid并存入数据库
	id, err := snowflake.CreateSnowID()
	if err != nil {
		return returnClientRegisterErr, logError, errSnowID
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
		return returnClientRegisterErr, logError, errInsert
	}

	return returnClientRegisterSuccess, logSuccess, returnClientRegisterSuccess
}

func LoginService(UserRe *user.Login) (string, string, string) {
	//1. 验证账户是否存在
	ok, err := sqluser.VerifyUserExits(UserRe.Username)
	if err != nil {
		return returnClientLoginErr, logError, errVerify
	}
	if !ok {
		return errLogin, logError, errLogin
	}

	//2. 存在此账户,则判断密码
	var sqlu user.Login
	if err := sqluser.VerifyUserLogin(UserRe.Username, &sqlu); err != nil {
		return returnClientLoginErr, logError, errSelect
	}

	if (UserRe.Username == sqlu.Username) && (UserRe.Password == sqlu.Password) {
		return returnClientLoginSuccess, logSuccess, returnClientLoginSuccess
	} else {
		return errPassword, logError, errPassword
	}
}
