package response

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

//const (
//	CodeSuccess = 200 + iota
//	CodeERROR
//)

const (
	CodeSuccess = 0
	CodeERROR   = -1
	CodeMax     = 100
)

type ResCode int64

var CodeMsg = map[ResCode]string{
	CodeSuccess: "操作成功",
	CodeERROR:   "操作失败",
	CodeMax:     "未知错误",
}

func (c ResCode) Msg() string {
	msg, ok := CodeMsg[c]
	if !ok {
		msg = CodeMsg[c]
	}
	return msg
}
