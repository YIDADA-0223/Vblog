package start

import (
	"os"

	"gitee.com/VBLOG/conf"
	"github.com/spf13/cobra"
)

var (
	testParam string
)
var Cmd = &cobra.Command{
	Use:   "start",
	Short: "start vblog",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		// 1. 加载配置
		configPath := os.Getenv("CONFIG_PATH")
		if configPath == "" {
			// configPath = "F:/Download/Go/PROJECT/VBLOG/etc/application.yml"
			configPath = "etc/application.yml"
		}

		// 4. 服务启动
		cobra.CheckErr(conf.C().Application.Start())
	},
}

func init() {
	// Flags()本地参数
	Cmd.Flags().StringVarP(&testParam, "test", "t", "test", "test flag")
}
