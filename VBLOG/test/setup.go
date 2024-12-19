package test

import (

	// 加载所有业务模块
	_ "gitee.com/VBLOG/apps"
	"gitee.com/VBLOG/conf"
	"gitee.com/VBLOG/ioc"
	"github.com/spf13/cobra"
)

func DevelopmentSetup() {
	// 加载配置,单元测试，通过环境变量读取的，vscode传递进来的
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}
	// 2. 初始化Ioc
	cobra.CheckErr(ioc.Controller.Init())
	// 3. 初始化Api
	cobra.CheckErr(ioc.Api.Init())
}
