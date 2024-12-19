package impl_test

import (
	"context"

	"gitee.com/VBLOG/apps/blog"
	"gitee.com/VBLOG/ioc"
	"gitee.com/VBLOG/test"

	// 导入被测试对象,最好是全部导入
	_ "gitee.com/VBLOG/apps"
)

var (
	//声明被测试的对象
	serviceImpl blog.Service
	ctx         = context.Background()
)

// 被测试的对象
func init() {
	test.DevelopmentSetup()
	// serviceImpl = impl.NewUserServiceImpl()
	serviceImpl = ioc.Controller.Get(blog.AppName).(blog.Service)
}
