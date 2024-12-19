package blog_test

import (
	"testing"

	"gitee.com/VBLOG/apps/blog"
)

func TestNewBlog(t *testing.T) {
	ins := blog.NewBlog()
	t.Log(ins.CreateBlogRequest)
}
