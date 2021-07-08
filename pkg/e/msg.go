package e

var MsgFlags = map[int]string{
	ERROR:        "fail",
	SUCCESS:      "ok",
	InvalidParam: "请求参数错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
