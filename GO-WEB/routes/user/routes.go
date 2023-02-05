package user

import (
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/go-web/controllers"
	"github.com/robot007num/go/go-web/utils"
)

var RegisterUserRoutes = func(e *gin.Engine) {

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)
	e.GET("/getsection", controllers.GetSection)
	e.GET("/getsection/:id", controllers.GetSectionClass)
	e.GET("/get-post/:classId", controllers.GetAllPost)
	e.GET("/get-post/:classId/:postId", controllers.GetSpecifyPost)

	v1 := e.Group("/v1").Use(utils.JWTAuthMiddleware())
	{
		v1.GET("/ping", controllers.Ping)
		v1.POST("/add-section", controllers.AddSection)
		v1.POST("/new-post", controllers.PostNewPost)
		v1.POST("/vote", controllers.UserVote)
	}

	//e.POST("/changePassword", controllers.ChangePassword)
}

// 点赞和点踩的思路：
/*  点赞1  取消0  点踩-1
{
User:
帖子ID:
行为：1/0/-1
}
1. 用redis里面的Set类型。因为Set 元素不重复且无序
2. 当用户点赞之后就写入到点赞集合
3. 当用户取消点赞后则从点赞集合去除
4. 当用户点踩之后写入到点踩集合
*/
