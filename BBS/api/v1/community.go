package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/bbs/dao"
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/model/common/request"
	"github.com/robot007num/go/bbs/model/common/response"
	"github.com/robot007num/go/bbs/utils"
	"go.uber.org/zap"
)

// GetAllCommunity
// @Tags     community
// @Summary  获取社区列表
// @Produce   application/json
// @Param    nil
// @Success  200   {object}  response.Response{data=systemRes.Community,msg=string}  "返回社区列表"
// @Router   /base/allCommunity [get]
func GetAllCommunity(c *gin.Context) {
	var all []response.Community
	if err := dao.SQLCommunityGet(&all); err != nil {
		utils.Result(-1, ErrSQL, c)
		global.GVA_LOG.Error("MYSQL数据库", zap.String("err:", err.Error()))
		return
	}

	utils.Result(0, all, c)
}

// AddCommunity
// @Tags     community
// @Summary  获取社区列表
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param    data  body      request.AddCommunity					true "社区,描述"
// @Success  200   {object}  response.Response{data=nil,msg=string}  "提示信息"
// @Router   /community/add [post]
func AddCommunity(c *gin.Context) {
	//1. 获取参数
	var u request.AddCommunity
	err := c.ShouldBindJSON(&u)
	if err != nil {
		utils.Result(-1, ErrorJsonBind, c)
		return
	}

	userid := utils.GetUserUuid(c)

	//2. 业务逻辑
	var su response.SQLUser
	su, err = dao.SQLUserSelectToUserID(userid)
	if err != nil {
		utils.Result(-1, ErrSQL, c)
		global.GVA_LOG.Error("MYSQL数据库", zap.String("err:", err.Error()))
		return
	}

	if su.Type != 2 {
		utils.Result(-1, ErrorUserAccess, c)
		return
	}

	ok, err := dao.SQLCommunityAdd(u)
	if err != nil {
		utils.Result(-1, ErrSQL, c)
		global.GVA_LOG.Error("MYSQL数据库", zap.String("err:", err.Error()))
		return
	}
	if !ok {
		utils.Result(-1, "已有相同的社区", c)
		return
	}

	utils.Result(0, nil, c)
}
