package response

import (
	"github.com/robot007num/go/bbs/model/common/internal"
	"reflect"
)

//SQLUser 存入数据库并在登录之后返回此结构
type SQLUser struct {
	internal.Basic
	UserId   int64  `json:"user_id" db:"user_id"`  //唯一ID
	Account  string `json:"account" db:"account" ` //登录账号
	Password string `json:"-" db:"password" `      //密码
	UserName string `json:"name" db:"username"`    //用户昵称
	Email    string `json:"email" db:"email"`      //邮件
	Type     int    `json:"type" db:"type"`        //用户类型 0普通 1版主 2超级
	Enable   int    `json:"enable" db:"enable"`    //用户是否被冻结 0正常 1冻结
}

type LoginResponse struct {
	User         SQLUser `json:"userinfo"`      //用户基本信息
	AccessToken  string  `json:"access_token"`  //访问Token
	RefreshToken string  `json:"refresh_token"` //刷新访问Token
}

//IsEmpty 判断SQL是否为空
func (a SQLUser) IsEmpty() bool {
	return reflect.DeepEqual(a, SQLUser{})
}
