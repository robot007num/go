package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/bbs/dao"
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/model/common/request"
	"github.com/robot007num/go/bbs/model/common/response"
	"github.com/robot007num/go/bbs/pkg"
	"github.com/robot007num/go/bbs/utils"
	"go.uber.org/zap"
	"strconv"
)

// PostGet
// @Tags     post
// @Summary  获取帖子详情
// @Produce   application/json
// @Param    nil
// @Success  200   {object}  response.Response{response.Post,msg=string}
// @Router   /post/:postId [get]
func PostGet(c *gin.Context) {
	id := c.Param("postId")
	fmt.Println(id)
	i64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.Result(-1, "string 转int64失败", c)
		global.GVA_LOG.Error("获取参数转换", zap.String("err:", err.Error()))
		return
	}

	pt, err := dao.PostGet(i64)
	if err != nil {
		utils.Result(-1, ErrSQL, c)
		global.GVA_LOG.Error("MYSQL数据库", zap.String("err:", err.Error()))
		return
	}

	utils.Result(0, pt, c)
}

// PostNew
// @Tags     post
// @Summary  新建帖子
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param    data  body      request.PostNew
// @Success  200   {object}  response.Response{response.Post,msg=string}
// @Router   /post/new [post]
func PostNew(c *gin.Context) {
	//1. 获取参数
	var u request.PostNew
	err := c.ShouldBindJSON(&u)
	if err != nil {
		utils.Result(-1, ErrorJsonBind, c)
		return
	}

	//2. 逻辑处理
	postId, _ := pkg.CreateSnowID()

	pt := response.Post{
		PostNew: request.PostNew{
			Title:     u.Title,
			Content:   u.Content,
			SectionId: u.SectionId,
		},
		PostId: postId,
		Userid: utils.GetUserUuid(c),
	}

	if err = dao.PostNew(pt); err != nil {
		utils.Result(-1, ErrSQL, c)
		global.GVA_LOG.Error("MYSQL数据库", zap.String("err:", err.Error()))
	}

	//3. 查询并返回
	pt, err = dao.PostGet(pt.PostId)
	if err != nil {
		utils.Result(-1, ErrSQL, c)
		global.GVA_LOG.Error("MYSQL数据库", zap.String("err:", err.Error()))
		return
	}

	utils.Result(0, pt, c)

}
