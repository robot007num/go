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

	//4. 存入Redis数据库
	//var userkey = "user:" + UserRe.Username
	//repository.GetRedisCon().Do("HSET", userkey, "access_token", allToken.AccessToken)
	//repository.GetRedisCon().Do("HSET", userkey, "refresh_toke", allToken.RefreshToken)

	return response.CodeLoginSuccess, logSuccess, ""
}

func VotePost(u user.VotePost) (response.ResCode, string, string) {
	//1. 当是0时,则代表要取消或者是点赞
	if u.VoteValue == 0 {
		// 从点赞去除

		// 从取消去除
	}

	//2. 1 则是代表要点赞
	{
		// 直接从点踩集合去掉
		// 加入点赞集合
	}

	// 3. -1 则是代表要点踩
	{
		// 直接从点赞集合去掉
		// 加入点踩集合
	}

	return response.CodeLoginSuccess, logSuccess, ""
}
