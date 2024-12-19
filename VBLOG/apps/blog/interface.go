package blog

import (
	"context"

	"gitee.com/VBLOG/common"
	"gitee.com/VBLOG/exception"
)

const (
	AppName = "blogs"
)

type Service interface {
	// 文章列表查询
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	// 文章详情
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	// 文章创建
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	// 文章更新
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	// 文章删除
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)
	// 文章发布
	UpdateBlogStatus(context.Context, *UpdateBlogStatusRequest) (*Blog, error)
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageRequest: common.NewPageRequest(),
	}
}

type QueryBlogRequest struct {
	*common.PageRequest
	// 关键字参数，模糊匹配，根据文章名称进行模糊搜索
	KeyWords string `json:"keywords"`
	// 状态过滤, 0/1, nil
	Status *Status `json:"status"`
}

func NewDescribeBlogRequest(id string) *DescribeBlogRequest {
	return &DescribeBlogRequest{
		BlogId: id,
	}
}

type DescribeBlogRequest struct {
	BlogId string `json:"blog_id"`
}

func NewUpdateBlogRequest(blogId string) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:            blogId,
		UpdateMode:        common.UPDATE_MODE_PUT,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}
func (req *UpdateBlogRequest) Validate() error {
	if req.CreateBlogRequest == nil {
		return exception.ErrValidateFailed("CreateBlogRequest required")
	}
	return common.Validate(req)
}

type UpdateBlogRequest struct {
	// 博客Id
	BlogId string `json:"blog_id" validate:"required"`
	// 更新模型 全量/部分更新
	UpdateMode common.UPDATE_MODE `json:"update_mode"`
	// 需要更新数据
	*CreateBlogRequest `validate:"required"`
}

func NewDeleteBlogRequest(id string) *DeleteBlogRequest {
	return &DeleteBlogRequest{
		BlogId: id,
	}
}

type DeleteBlogRequest struct {
	BlogId string `json:"blog_id"`
}

func NewUpdateBlogStatusRequest(blogid string) *UpdateBlogStatusRequest {
	return &UpdateBlogStatusRequest{
		BlogId:                   blogid,
		ChangedBlogStatusRequest: &ChangedBlogStatusRequest{},
	}
}

type UpdateBlogStatusRequest struct {
	BlogId string `json:"blog_id"`
	*ChangedBlogStatusRequest
}
