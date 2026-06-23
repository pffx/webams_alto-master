package e

var msgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	ERROR_ACTION_FAILED:            "操作失败",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "The username or password are incorrect",
	ERROR_PERMISSION:               "权限失败",
	ERROR_SSH_CREATE_FAIL:          "SSH 连接失败",
	ERROR_TEMPLATE_CREATE_FAIL:     "Template 创建失败",
	ERROR_FIND_OLTS_FAIL:           "OLT 读取失败",
	ERROR_NO_FILE:                  "没有找到文件",
}

// GetMessage 通过code获取message
func GetMessage(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}

	return msgFlags[ERROR]
}
