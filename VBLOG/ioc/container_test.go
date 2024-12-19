package ioc_test

import (
	"testing"

	user "gitee.com/VBLOG/apps/user/impl"
	"gitee.com/VBLOG/ioc"
)

func TestRegistry(t *testing.T) {
	ioc.Controller.Registry("user", &user.UserServiceImpl{})
	t.Logf("%p", ioc.Controller.Get("user"))
}
