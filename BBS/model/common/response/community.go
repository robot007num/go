package response

import (
	"github.com/robot007num/go/bbs/model/common/internal"
	"github.com/robot007num/go/bbs/model/common/request"
)

type Community struct {
	internal.Basic
	//SectionName  string `json:"section_name" db:"section_name" binding:"required"` //社区名字
	//Introduction string `json:"introduction" db:"introduction" binding:"required"` //简单描述
	request.AddCommunity
}
