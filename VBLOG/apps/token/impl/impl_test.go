package impl_test

import (
	"context"
	"testing"

	"gitee.com/VBLOG/apps/token"
	"gitee.com/VBLOG/ioc"
	"gitee.com/VBLOG/test"
	// user "gitee.com/VBLOG/apps/user/impl"
)

var (
	//声明被测试的对象
	serviceImpl token.Service
	ctx         = context.Background()
)

// 招对象
func init() {
	//使用构造函数
	// serviceImpl = impl.NewTokenServiceImpl(user.NewUserServiceImpl())
	// 初始化单测环境
	test.DevelopmentSetup()
	// 去ioc中获取被删除的业务对象
	serviceImpl = ioc.Controller.Get(token.AppName).(token.Service)
}

// 颁发Token测试
func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest("admin", "123456")
	tk, err := serviceImpl.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk.String())
}

// 撤销Token是需要access_token与refresh_token
func TestRevolkToken(t *testing.T) {
	req := token.NewRevolkTokenRequest("cr4p4mols52oed3rh30g", "cr4p4mols52oed3rh310")
	tk, err := serviceImpl.RevolkToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

// 校验Token测试
func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("cr4sh6gls52oc836lmn0")
	tk, err := serviceImpl.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}
