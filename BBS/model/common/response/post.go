package response

import (
	"github.com/robot007num/go/bbs/model/common/internal"
	"github.com/robot007num/go/bbs/model/common/request"
)

type Post struct {
	internal.Basic
	request.PostNew
	PostId       int64 `json:"post_id" db:"post_id" binding:"required" `             //帖子ID
	Userid       int64 `json:"user_id" db:"user_id" binding:"required" `             //创建人
	Type         int   `json:"type" db:"type" binding:"required" `                   //帖子类型 0:普通(default) 1:置顶
	Status       int   `json:"status" db:"status" binding:"required" `               //帖子状态 0:正常(default) 1:精华 2:黑名单
	CommentCount int   `json:"comment_count" db:"comment_count" binding:"required" ` //评论数量
	Score        int   `json:"score" db:"score" binding:"required" `                 //帖子热度(用于按照热度排行帖子)
}
