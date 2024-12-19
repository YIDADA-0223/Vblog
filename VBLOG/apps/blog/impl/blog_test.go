package impl_test

import (
	"testing"

	"gitee.com/VBLOG/apps/blog"
	"gitee.com/VBLOG/common"
)

func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	req.Title = "Go 全栈开发"
	req.Author = "author"
	req.Content = "Md内容填充"
	req.Summary = "文章概要信息"
	ins, err := serviceImpl.CreateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
func TestQueryBlog(t *testing.T) {
	req := blog.NewQueryBlogRequest()
	ins, err := serviceImpl.QueryBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
func TestDescribeBlog(t *testing.T) {
	req := blog.NewDescribeBlogRequest("1")
	ins, err := serviceImpl.DescribeBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
func TestPathUpdateBlog(t *testing.T) {
	req := blog.NewUpdateBlogRequest("1")
	req.UpdateMode = common.UPDATE_MODE_PATCH
	req.Title = "更新后文章标题1"
	ins, err := serviceImpl.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
func TestPutUpdateBlog(t *testing.T) {
	req := blog.NewUpdateBlogRequest("1")
	req.UpdateMode = common.UPDATE_MODE_PUT
	req.Title = "更新后文章标题PUT"
	req.Author = "patch"
	req.Content = "patch"
	ins, err := serviceImpl.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
func TestUpdateBlogStatus(t *testing.T) {
	req := blog.NewUpdateBlogStatusRequest("1")
	req.SetStatus(blog.STATUS_PUBLISH)
	ins, err := serviceImpl.UpdateBlogStatus(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
func TestDeleteBlog(t *testing.T) {
	req := blog.NewDeleteBlogRequest("1")
	ins, err := serviceImpl.DeleteBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
