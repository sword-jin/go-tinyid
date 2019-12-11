package constant

type ErrorCode struct {
	Code    int
	Message string
}

var BizTypeErr ErrorCode
var TokenErr ErrorCode
var SysErr ErrorCode

func init() {
	BizTypeErr = ErrorCode{
		Code:    5,
		Message: "bizType error",
	}
	TokenErr = ErrorCode{
		Code:    5,
		Message: "token is error",
	}
	SysErr = ErrorCode{
		Code:    6,
		Message: "sys error",
	}
}
