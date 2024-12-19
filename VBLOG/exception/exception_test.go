package exception_test

import (
	"testing"

	"gitee.com/VBLOG/exception"
)

func CheckIsError() error {
	return exception.NewApiExcepiton(50001, "用户名或者密码不正确")
}

func TestIsError(t *testing.T) {
	err := CheckIsError()
	t.Log(err)
	//怎么获取ErrorCode，断言这个接口的对象的具体类型
	if v, ok := err.(*exception.ApiExcepiton); ok {
		t.Log(v.Code)
		t.Log(v.String())
	}
	//前端想要获取的是一个完整ApiExcepiton，该怎么获取
}
