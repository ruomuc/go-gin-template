package e

var MsgFlags = map[int]string{
	ERROR:        "fail",
	SUCCESS:      "ok",
	InvalidParam: "请求参数错误",

	ErrorPassword: "用户名密码错误",

	AuthorizationnNotFound: "authentication参数不能为空",
	AuthTokenExpired:       "token已过期",
	AuthTokenFailed:        "token校验失败",
	TokenGenerateFailed:    "token生成失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
