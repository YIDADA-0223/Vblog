package main

import (

	// 注册所有的业务模块
	_ "gitee.com/VBLOG/apps"
	"gitee.com/VBLOG/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}

}
