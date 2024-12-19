package api

import (
	"gitee.com/VBLOG/apps/blog"
	"gitee.com/VBLOG/conf"
	"gitee.com/VBLOG/ioc"
)

func NewTokenApiHandler() *BlogApiHandler {
	return &BlogApiHandler{
		svc: ioc.Controller.Get(blog.AppName).(blog.Service),
	}
}

// 注册api
func init() {
	ioc.Api.Registry(blog.AppName, &BlogApiHandler{})
}

// 核心处理APi请求
type BlogApiHandler struct {
	// 依赖TokenServiceImpl
	// impl.TokenServiceImpl 不推荐这样使用
	//依赖接口
	svc blog.Service
}

// 自己决定初始化逻辑
func (h *BlogApiHandler) Init() error {
	h.svc = ioc.Controller.Get(blog.AppName).(blog.Service)
	// 已经前置在config中/vblog/v1/
	subRouter := conf.C().Application.GinRootRouter().Group("blogs")
	h.Registry(subRouter)
	return nil
}
