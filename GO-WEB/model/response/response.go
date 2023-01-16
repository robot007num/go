package response

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
	CodeMax
)

var CodeMsg = map[ResCode]string{
	CodeRegisterSuccess:   "注册成功",
	CodeRegisterError:     "注册失败",
	CodeLoginSuccess:      "登录成功",
	CodeMax:               "新的错误",
	CodeLoginError:        "登录失败",
	CodeInvalidParameters: "参数错误",
}

const (
	InfoUserRegister   = "该用户已注册"
	InfoUserVerify     = "验证用户失败"
	InfoUserSnowID     = "生成UserID失败"
	InfoUserInsert     = "插入新用户失败"
	InfoUserUnRegister = "该用户未注册"
	InfoUserSelect     = "查询用户失败"
	InfoUserPassword   = "用户名/密码错误"
)

func (c ResCode) Msg() string {
	msg, ok := CodeMsg[c]
	if !ok {
		msg = CodeMsg[c]
	}
	return msg
}
