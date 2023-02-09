package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/model/common/request"
	"github.com/robot007num/go/bbs/model/common/response"
	"github.com/robot007num/go/bbs/utils"
	"go.uber.org/zap"
	"strings"
)

const (
	ErrorJsonBind = "传参有误"
	ErrorTokenNil = "Refresh Token为空"
)

// Login
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      systemReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户ID和Name,双token"
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
	su.Account = "test01"
	su.UserId = "0"

	//3. 生成Token
	A, R, e := TokenNext(c, su)
	if e != nil {
		utils.Result(-1, err, c)
		return
	}

	r := response.LoginResponse{
		User: response.SQLUser{
			UserId:  su.UserId,
			Account: su.Account,
		},
		AccessToken:  A,
		RefreshToken: R,
	}

	//这里用自定义规则生成有"状态"的Token
	//自定义规则：将格式为JWT的refresh_token转变为一个仅作为匹配一致性使用的、无意义的字符串,
	//可以对整个refresh_token进行一次md5计算（长度32），也可以直接把jwt token的签名部分截取使用（长度43）

	//获取RefreshToken的签名
	tokenValue := strings.Split(R, ".")
	fmt.Println("这是Login存入redis的token:", tokenValue[2])
	//并将最新的RefreshToken存入(相当于删除)
	if err = utils.SetRedisJWT(tokenValue[2], su.UserId); err != nil {
		global.GVA_LOG.Info("redis存Token值失败")
	}

	//4. 返客户端
	utils.Result(0, r, c)
}

// TokenNext 登录以后签发jwt
func TokenNext(c *gin.Context, user response.SQLUser) (a, r string, err error) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	Acclaims, Reclaims := j.CreateClaims(request.BaseClaims{
		UserID:   user.UserId,
		UserName: user.Name,
	})

	Atoken, err := j.CreateToken(Acclaims)
	if err != nil {
		global.GVA_LOG.Error("获取Atoken失败!", zap.Error(err))
		return "", "", errors.New("获取Atoken失败!")
	}

	Rtoken, err := j.CreateToken(Reclaims)
	if err != nil {
		global.GVA_LOG.Error("获取Rtoken失败!", zap.Error(err))
		return "", "", errors.New("获取Rtoken失败!")
	}

	return Atoken, Rtoken, nil
}

func Register(c *gin.Context) {
	//1. 获取参数
	var u request.Login
	err := c.ShouldBindJSON(&u)
	if err != nil {
		utils.Result(-1, ErrorJsonBind, c)
		return
	}
	//2. 业务逻辑

	//3. 返客户端
	utils.Result(0, nil, c)
}

//RefreshToken 刷新过期的Access Token
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
	redisToken, err = utils.GetRedisJWT(claims.UserID)
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
	if err = utils.SetRedisJWT(tokenValue[2], claims.UserID); err != nil {
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

func ChangePassword(c *gin.Context) {
	utils.Result(0, "检验Token成功", c)
}
