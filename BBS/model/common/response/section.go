package response

import (
	"github.com/robot007num/go/bbs/model/common/internal"
	"github.com/robot007num/go/bbs/model/common/request"
)

type GetSection struct {
	internal.Basic
	request.AddSection
}
