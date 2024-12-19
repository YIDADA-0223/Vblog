package exception

import "encoding/json"

func NewApiExcepiton(code int, message string) *ApiExcepiton {
	return &ApiExcepiton{
		Code:    code,
		Message: message,
	}
}

// 用于描述业务异常
type ApiExcepiton struct {
	//业务异常的编码，50001 表示Token过期
	Code int `json:"code"`
	// 异常描述信息
	Message string `json:"message"`
	// Http状态码,不会出现在Boyd里面，序列化json是看不到的
	HttpCode int `json:"-"`
}

func (e *ApiExcepiton) Error() string {
	return e.Message
}
func (e *ApiExcepiton) String() string {
	dj, _ := json.MarshalIndent(e, "", " ")
	return string(dj)
}
func (e *ApiExcepiton) WithMessage(msg string) *ApiExcepiton {
	e.Message = msg
	return e
}
func (e *ApiExcepiton) WithHttpCode(httpcode int) *ApiExcepiton {
	e.HttpCode = httpcode
	return e
}
