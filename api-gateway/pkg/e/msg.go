package e

var MsgFlags = map[int32]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	ErrorUploadFail:            "文件上传失败",
	ErrorDownloadFail:          "文件保存失败",
	ErrorSnapshotFail:          "缩略图生成失败",
	InvalidParams:              "请求参数错误",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
	ErrorNotCompare:            "不匹配",
	ErrorDatabase:              "数据库操作出错,请重试",
	ErrorAuthNotFound:          "Token不能为空",

	ErrorServiceUnavailable: "过载保护，服务暂时不可用",
	ErrorDeadline:           "服务调用超时",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int32) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
