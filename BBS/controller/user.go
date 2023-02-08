package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/model/common/request"
	"github.com/robot007num/go/bbs/model/common/response"
	"github.com/robot007num/go/bbs/utils"
	"go.uber.org/zap"
)

const (
	ErrorJsonBind = "传参有误"
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

func ChangePassword(c *gin.Context) {
	utils.Result(0, "检验Token成功", c)
}
