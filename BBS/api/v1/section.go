package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/bbs/dao"
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/model/common/request"
	"github.com/robot007num/go/bbs/model/common/response"
	"github.com/robot007num/go/bbs/utils"
	"go.uber.org/zap"
	"strconv"
)

// GetSection
// @Tags     section
// @Summary  获取从属社区的部门
// @Produce  application/json
// @Param    nil
// @Success  200   {object}  response.Response{data=[]response.GetSection,msg=string}
// @Router   /base/section/:communityID [get]
func GetSection(c *gin.Context) {
	//获取参数
	id, _ := strconv.Atoi(c.Param("communityID"))
	var all []response.GetSection
	if err := dao.SQLSectionGet(&all, id); err != nil {
		utils.Result(-1, ErrSQL, c)
		return
	}
	if len(all) == 0 {
		utils.Result(0, nil, c)
		return
	}

	utils.Result(0, all, c)
}

// AddSection
// @Tags     section
// @Summary  增加部门
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param    data  body      request.AddSection
// @Success  200   {object}  response.Response{msg=string}
// @Router   /section/add [post]
func AddSection(c *gin.Context) {
	//1. 获取参数
	var u request.AddSection
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

	ok, err := dao.SQLSectionAdd(u)
	if err != nil {
		utils.Result(-1, ErrSQL, c)
		global.GVA_LOG.Error("MYSQL数据库", zap.String("err:", err.Error()))
		return
	}
	if !ok {
		utils.Result(-1, "已有相同的部门", c)
		return
	}

	utils.Result(0, nil, c)
}

// GetSectionPost
// @Tags     section
// @Summary  获取部门所属的所有帖子
// @Produce  application/json
// @Param    nil
// @Success  200   {object}  response.Response{data=[]response.GetSection,msg=string}
// @Router   /base/section/:communityID/post [get]
func GetSectionPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("communityID"))

	p, err := dao.PostGetAll(id)
	if err != nil {
		utils.Result(-1, ErrSQL, c)
		global.GVA_LOG.Error("MYSQL数据库", zap.String("err:", err.Error()))
		return
	}
	utils.Result(0, p, c)
}
