package impl

import (
	"gitee.com/VBLOG/apps/blog"
	"gitee.com/VBLOG/conf"
	"gitee.com/VBLOG/ioc"
	"gorm.io/gorm"
)

func init() {
	ioc.Controller.Registry(blog.AppName, &BlogServiceImpl{})
}

type BlogServiceImpl struct {
	db *gorm.DB
}

func (i *BlogServiceImpl) Init() error {
	i.db = conf.C().MySQL.GetDB()
	return nil
}
