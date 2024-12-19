package cmd

import (
	"fmt"
	// 加载所有业务模块
	_ "gitee.com/VBLOG/apps"
	initCmd "gitee.com/VBLOG/cmd/init"
	"gitee.com/VBLOG/cmd/start"
	"gitee.com/VBLOG/conf"
	"gitee.com/VBLOG/ioc"
	"github.com/spf13/cobra"
)

var (
	configPath string
)
var RootCmd = &cobra.Command{
	Use:   "vblog",
	Short: "vblog service",
	Run: func(cmd *cobra.Command, args []string) {

		// Do Stuff Here
		if len(args) > 0 {
			if args[0] == "version" {
				fmt.Println("v0.0.1")
			}
		} else {
			cmd.Help()
		}

	},
}

func Execute() error {
	// 初始化需要执行的逻辑
	cobra.OnInitialize(func() {
		// 1. 加载配置
		cobra.CheckErr(conf.LoadConfigFromYaml(configPath))
		// 2. 初始化Ioc
		cobra.CheckErr(ioc.Controller.Init())
		// 3. 初始化Api
		cobra.CheckErr(ioc.Api.Init())
	})
	return RootCmd.Execute()
}
func init() {
	//PersistentFlags() 全局标签
	RootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "etc/application.yml", "the service config file")
	RootCmd.AddCommand(initCmd.Cmd)
	RootCmd.AddCommand(start.Cmd)
}
