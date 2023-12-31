package controller

// ResCode 状态码
type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeInvalidToken
	CodeInvalidAuthFormat
	CodeNeedLogin
	ErrVoteRepeated
	ErrorVoteTimeExpire
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeNeedLogin:         "需要登录",
	CodeInvalidToken:      "无效的token",
	CodeInvalidAuthFormat: "认证格式有误",
	ErrVoteRepeated:       "重复投票",
	ErrorVoteTimeExpire:   "投票时间已过",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
