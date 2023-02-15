package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/bbs/dao"
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/model/common/request"
	"github.com/robot007num/go/bbs/model/common/response"
	"github.com/robot007num/go/bbs/pkg"
	"github.com/robot007num/go/bbs/utils"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

const (
	ErrorJsonBind           = "传参有误"
	ErrorTokenNil           = "Refresh Token为空"
	ErrorUserExits          = "用户已存在"
	ErrorUserNotExits       = "用户不存在"
	ErrorUserPassword       = "原密码不正确"
	ErrorUserPasswordChange = "修改密码失败"
	ErrorUserAccess         = "权限错误"
	ErrSQL                  = "数据库操作失败"
)

// TokenNext
// @Summary  用户登录以后签发jwt
func TokenNext(c *gin.Context, user response.SQLUser) (a, r string, err error) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	Acclaims, Reclaims := j.CreateClaims(request.BaseClaims{
		UserID:   user.UserId,
		UserName: user.UserName,
	})

	accessToken, err := j.CreateToken(Acclaims)
	if err != nil {
		global.GVA_LOG.Error("获取Access_token失败!", zap.Error(err))
		return "", "", errors.New("获取Access_token失败")
	}

	refreshToken, err := j.CreateToken(Reclaims)
	if err != nil {
		global.GVA_LOG.Error("获取Refresh_token失败!", zap.Error(err))
		return "", "", errors.New("获取Refresh_token失败")
	}

	return accessToken, refreshToken, nil
}

//RefreshToken
//@Summary 刷新过期的Access Token
func RefreshToken(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	if token == "" {
		utils.Result(-1, ErrorTokenNil, c)
		c.Abort()
		return
	}

	//2. 解析Token
	//上述操作中的第二、三种验证如果出现任何原因的续签失败，
	//则统一回到1步骤要求用户重新登陆。（视情况也可增加盗号风险提醒，后详）
	j := utils.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		utils.Result(-99, nil, c)
		c.Abort()
		return
	}

	//获取RefreshToken的签名
	tokenValue := strings.Split(token, ".")

	//保证RefreshToken 只能一次性使用
	var redisToken string
	redisToken, err = utils.GetRedisJWT(strconv.FormatInt(int64(claims.UserID), 10))
	if err != nil {
		global.GVA_LOG.Info("根据userid获取redis的value值失败")
	}
	if redisToken != tokenValue[2] {
		global.GVA_LOG.Info("与最新存入redis的RefreshToken不符合")
		utils.Result(-1, "重复使用", c)
		c.Abort()
		return
	}

	//3. 当RefreshToken 有效时 重新申请一对拥有全新过期时间的长短令牌
	var su response.SQLUser
	su.Account = claims.UserName
	su.UserId = claims.UserID

	A, R, e := TokenNext(c, su)
	if e != nil {
		utils.Result(-1, err, c)
		return
	}

	//获取RefreshToken的签名
	tokenValue = strings.Split(R, ".")
	//并将最新的RefreshToken存入(相当于删除)
	if err = utils.SetRedisJWT(tokenValue[2], strconv.FormatInt(int64(claims.UserID), 10)); err != nil {
		global.GVA_LOG.Info("redis存Token值失败")
	}

	r := response.LoginResponse{
		User: response.SQLUser{
			UserId:  su.UserId,
			Account: su.Account,
		},
		AccessToken:  A,
		RefreshToken: R,
	}
	//4. 返客户端
	utils.Result(0, r, c)

}

// Login
// @Tags     user
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      request.Login                                            true  "账户, 密码"
// @Success  200   {object}  response.Response{data=response.LoginResponse,msg=string}  "返回用户所有基本信息,双token"
// @Router   /base/login [post]
func Login(c *gin.Context) {
	//1. 获取参数
	var u request.Login
	err := c.ShouldBindJSON(&u)
	if err != nil {
		utils.Result(-1, ErrorJsonBind, c)
		return
	}
	//2. 校验用户名和密码是否正确
	var su response.SQLUser
	su, err = dao.SQLUserSelectToAccount(u.Account)
	if err != nil {
		utils.Result(-1, ErrorUserNotExits, c)
		return
	}

	//3. 生成Token
	A, R, e := TokenNext(c, su)
	if e != nil {
		utils.Result(-1, err, c)
		return
	}

	r := response.LoginResponse{
		User:         su,
		AccessToken:  A,
		RefreshToken: R,
	}

	//这里用自定义规则生成有"状态"的Token 获取RefreshToken的签名
	tokenValue := strings.Split(R, ".")
	//fmt.Println("这是Login存入redis的token:", tokenValue[2])
	//并将最新的RefreshToken存入(相当于删除)
	if err = utils.SetRedisJWT(tokenValue[2], strconv.FormatInt(int64(su.UserId), 10)); err != nil {
		global.GVA_LOG.Info("redis存Token值失败")
	}

	//4. 返客户端
	utils.Result(0, r, c)
}

// Register
// @Tags     user
// @Summary  用户注册
// @Produce   application/json
// @Param    data  body      request.Login                                            true  "用户名, 密码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回用户所有基本信息"
// @Router   /base/Register [post]
func Register(c *gin.Context) {
	//1. 获取参数
	var u request.Login
	err := c.ShouldBindJSON(&u)
	if err != nil {
		utils.Result(-1, ErrorJsonBind, c)
		return
	}

	//2. 业务逻辑
	//验证账户是否已注册
	_, err = dao.SQLUserSelectToAccount(u.Account)
	if err == nil {
		utils.Result(-1, ErrorUserExits, c)
		return
	}

	userid, _ := pkg.CreateSnowID()
	uuid, _ := pkg.CreateUUID()
	username := "BBS用户" + strconv.FormatInt(int64(uuid), 10) //生成随机用户名

	sqlU := response.SQLUser{UserId: userid, Account: u.Account, Password: u.Password, UserName: username}

	err = dao.SQLInsertUser(sqlU)
	if err != nil {
		utils.Result(-1, ErrSQL, c)
		global.GVA_LOG.Error("MYSQL数据库", zap.String("err:", err.Error()))
		return
	}

	sqlU, err = dao.SQLUserSelectToAccount(u.Account)
	if err != nil {
		utils.Result(-1, ErrSQL, c)
		global.GVA_LOG.Error("MYSQL数据库", zap.String("err:", err.Error()))
		return
	}

	//3. 返客户端
	global.GVA_LOG.Info("注册账户", zap.String("status:", "成功"))
	utils.Result(0, sqlU, c)
}

// ChangePassword
// @Tags     user
// @Summary  用户修改密码
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param    data  body      request.ChangePasswordReq      true  "旧密码,新密码"
// @Success  200   {object}  response.Response{msg=string}  "修改密码成功"
// @Router   /user/ChangePassword [post]
func ChangePassword(c *gin.Context) {
	//1. 检验参数
	var u request.ChangePasswordReq
	err := c.ShouldBindJSON(&u)
	if err != nil {
		utils.Result(-1, ErrorJsonBind, c)
		return
	}

	userid := utils.GetUserUuid(c)

	//2. 业务逻辑
	var old string
	old, err = dao.SQLUserSelectPart(userid)
	if err != nil {
		utils.Result(-1, ErrSQL, c)
		global.GVA_LOG.Error("MYSQL数据库", zap.String("err:", err.Error()))
		return
	}

	if old != u.Password {
		utils.Result(-1, ErrorUserPassword, c)
		return
	}

	err = dao.SQLUserChange(u.NewPassword, userid)
	if err != nil {
		utils.Result(-1, ErrorUserPasswordChange, c)
		return
	}

	//3. 返回
	global.GVA_LOG.Info("用户修改密码", zap.String("status", "成功"))
	utils.Result(0, "修改密码成功", c)
}
