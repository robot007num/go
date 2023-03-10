package response

import "net/http"

/* 返回给客户端的结构体
{
	code: //状态码
	msg: //信息
	Data: //数据
}
*/

type ReturnData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"message"`
	Data interface{} `json:"data"`
}

type ResCode int64

const (
	CodeRegisterSuccess = 200 + iota
	CodeLoginSuccess
	CodeRegisterError
	CodeLoginError
	CodeInvalidParameters
	CodeInsertSectionSuccess
	CodeInsertSectionError
	CodeGetSectionSuccess
	CodeGetSectionError
	CodeNewPostSuccess
	CodeNewPostError
	CodeGetPostSuccess
	CodeGetPostError
	CodeMax
)

var CodeMsg = map[ResCode]string{
	CodeRegisterSuccess:      "注册成功",
	CodeRegisterError:        "注册失败",
	CodeLoginSuccess:         "登录成功",
	CodeMax:                  "新的错误",
	CodeLoginError:           "登录失败",
	CodeInvalidParameters:    "参数错误",
	CodeInsertSectionSuccess: "插入板块成功",
	CodeInsertSectionError:   "插入板块失败",
	CodeGetSectionSuccess:    "获取板块成功",
	CodeGetSectionError:      "获取板块失败",
	CodeNewPostSuccess:       "帖子新建成功",
	CodeNewPostError:         "帖子新建失败",
	CodeGetPostSuccess:       "获取帖子成功",
	CodeGetPostError:         "获取帖子失败",
}

//记录错误详细信息
const (
	InfoUserRegister   = "该用户已注册"
	InfoUserVerify     = "验证用户失败"
	InfoUserSnowID     = "生成UserID失败"
	InfoUserInsert     = "插入新用户失败"
	InfoUserUnRegister = "该用户未注册"
	InfoUserSelect     = "查询用户失败"
	InfoUserPassword   = "用户名/密码错误"
	InfoTokenCreate    = "生成Token失败"
	InfoTokenNothing   = "请求头中Token为空"
	InfoTokenFormat    = "请求头中Token格式不正确"
	InfoTokenInvalid   = "请求头中Token失效"
	InfoTokenNumber    = "请求头中缺失AccessToken或者RefreshToken"
	InfoSectionSame    = "已有相同的板块"
	InfoSectionFail    = "板块查询失败"
)

var (
	HttpOK = http.StatusOK
)

func (c ResCode) Msg() string {
	msg, ok := CodeMsg[c]
	if !ok {
		msg = CodeMsg[c]
	}
	return msg
}
