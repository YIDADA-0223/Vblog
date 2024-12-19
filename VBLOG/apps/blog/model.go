package blog

import (
	"encoding/json"
	"time"

	"gitee.com/VBLOG/common"
)

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

type BlogSet struct {
	Total int64   `json:"total"`
	Items []*Blog `json:"item"`
}

func (b *BlogSet) String() string {
	dj, _ := json.MarshalIndent(b, "", " ")
	return string(dj)
}
func NewBlog() *Blog {
	return &Blog{
		common.NewMeta(),
		&CreateBlogRequest{
			Tags: map[string]string{},
		},
		&ChangedBlogStatusRequest{
			Status: STATUS_DRAFT,
		},
	}
}

type Blog struct {
	*common.Meta
	*CreateBlogRequest
	*ChangedBlogStatusRequest
}

func (b *Blog) String() string {
	dj, _ := json.MarshalIndent(b, "", " ")
	return string(dj)
}
func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags: map[string]string{},
	}
}

type CreateBlogRequest struct {
	// 文章标题
	Title string `json:"title" gorm:"column:title" validate:"required"`
	// 作者
	Author string `json:"author" gorm:"column:author" validate:"required"`
	// 文章内容
	Content string `json:"content" gorm:"column:content" validate:"required"`
	// 文章概要信息
	Summary string `json:"summary" gorm:"column:summary"`
	// 创建人
	CreateBy string `json:"create_by" gorm:"column:create_by"`
	// 标签 https://gorm.io/docs/serializer.html
	Tags map[string]string `json:"tags" gorm:"column:tags;serializer:json"`
}

func (req *CreateBlogRequest) Validate() error {
	return common.Validate(req)
}
func (c *CreateBlogRequest) String() string {
	dj, _ := json.MarshalIndent(c, "", " ")
	return string(dj)
}
func (req *ChangedBlogStatusRequest) SetStatus(s Status) {
	req.Status = s
	switch req.Status {
	case STATUS_PUBLISH:
		req.PublishedAt = time.Now().Unix()
	}
}

type ChangedBlogStatusRequest struct {
	// 发布时间
	PublishedAt int64 `json:"published_at" gorm:"column:published_at"`
	// 文章状态: 草稿/已发布
	Status Status `json:"status" gorm:"column:status"`
}

func (c *ChangedBlogStatusRequest) String() string {
	dj, _ := json.MarshalIndent(c, "", " ")
	return string(dj)
}
