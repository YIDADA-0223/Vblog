package impl

import (
	"context"

	"dario.cat/mergo"
	"gitee.com/VBLOG/apps/blog"
	"gitee.com/VBLOG/common"
	"gitee.com/VBLOG/exception"
)

// 文章列表查询
func (i *BlogServiceImpl) QueryBlog(ctx context.Context, in *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	set := blog.NewBlogSet()
	// 1.因为有默认值，不需要用户传参数的
	// 2.直接查询数据库
	query := i.db.WithContext(ctx).Table("blogs")
	if in.KeyWords != "" {
		query = query.Where("title like ?", "%"+in.KeyWords+"%")
	}
	if in.Status != nil {
		query = query.Where("status = ?", *in.Status)
	}
	// Count总数统计
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}
	// 查询
	err = query.Order("created_at DESC").Limit(in.PageSize).Offset(in.Offset()).Find(&set.Items).Error
	if err != nil {
		return nil, err
	}
	// err = query.Limit(in.PageSize).Offset(in.Offset()).Find(&set.Items).Error
	// if err != nil {
	// 	return nil, err
	// }

	return set, nil
}

// 文章详情
func (i *BlogServiceImpl) DescribeBlog(ctx context.Context, in *blog.DescribeBlogRequest) (*blog.Blog, error) {

	ins := blog.NewBlog()
	err := i.db.WithContext(ctx).Where("id = ?", in.BlogId).First(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 文章创建
func (i *BlogServiceImpl) CreateBlog(ctx context.Context, in *blog.CreateBlogRequest) (*blog.Blog, error) {
	// 1.验证请求参数
	if err := in.Validate(); err != nil {
		return nil, exception.ErrValidateFailed(err.Error())
	}
	// 2.构造实例对象
	ins := blog.NewBlog()
	ins.CreateBlogRequest = in
	// 3.入库返回
	//INSERT INTO `blogs` (`created_at`,`updated_at`,`title`,`author`,`content`,`summary`,`create_by`,`tags`,`published_at`,`status`) VALUES (1719648679,0,'Go 全栈开发','QUYI','Md内容填充','文章概要信息','','{}',0,0)
	err := i.db.WithContext(ctx).Create(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 文章更新
func (i *BlogServiceImpl) UpdateBlog(ctx context.Context, in *blog.UpdateBlogRequest) (*blog.Blog, error) {
	// 1. 先把需要更新的对象查询处理
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(in.BlogId))
	if err != nil {
		return nil, err
	}
	// 2.req，update
	switch in.UpdateMode {
	// 全量更新,你传什么就保存什么
	case common.UPDATE_MODE_PUT:
		ins.CreateBlogRequest = in.CreateBlogRequest
	// 增量更新，只更新有变化的字段
	case common.UPDATE_MODE_PATCH:
		if in.Author != "" {
			ins.Author = in.Author
		}
		if in.Content != "" {
			ins.Content = in.Content
		}
		err := mergo.MapWithOverwrite(ins.CreateBlogRequest, in.CreateBlogRequest)
		if err != nil {
			return nil, err
		}
	}
	// 校验更新的字段
	if err := ins.CreateBlogRequest.Validate(); err != nil {
		return nil, exception.ErrValidateFailed(err.Error())
	}
	// 执行更新,增量
	err = i.db.WithContext(ctx).Table("blogs").Save(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 文章删除
func (i *BlogServiceImpl) DeleteBlog(ctx context.Context, in *blog.DeleteBlogRequest) (*blog.Blog, error) {
	// 1. 先把需要更新的对象查询处理
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(in.BlogId))
	if err != nil {
		return nil, err
	}
	err = i.db.WithContext(ctx).Table("blogs").Where("id = ?", in.BlogId).Delete(&blog.Blog{}).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 文章发布
func (i *BlogServiceImpl) UpdateBlogStatus(ctx context.Context, in *blog.UpdateBlogStatusRequest) (*blog.Blog, error) {
	// 1. 先把需要更新的对象查询处理
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(in.BlogId))
	if err != nil {
		return nil, err
	}
	// 更新指定字段
	ins.ChangedBlogStatusRequest = in.ChangedBlogStatusRequest
	ins.SetStatus(in.Status)
	err = i.db.WithContext(ctx).Table("blogs").Where("id = ?", in.BlogId).Updates(ins.ChangedBlogStatusRequest).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}
