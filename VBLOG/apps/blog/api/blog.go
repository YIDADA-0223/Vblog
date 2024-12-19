package api

import (
	"gitee.com/VBLOG/apps/blog"
	"gitee.com/VBLOG/apps/token"
	"gitee.com/VBLOG/apps/user"
	"gitee.com/VBLOG/common"
	"gitee.com/VBLOG/exception"
	"gitee.com/VBLOG/middleware"
	"gitee.com/VBLOG/response"
	"github.com/gin-gonic/gin"
)

// 把自己的路由信息注册给Gin Root Router
// 每个业务模块，有每个业务模块的子路由
func (h *BlogApiHandler) Registry(appRouter gin.IRouter) {
	appRouter.GET("/", h.QueryBlog)
	appRouter.Use(middleware.Auth)
	appRouter.GET("/:id", h.DescribeBlog)
	appRouter.POST("/", middleware.RequireRole(user.ROLE_ADMIN), h.CreateBlog)
	appRouter.PUT("/:id", middleware.RequireRole(user.ROLE_ADMIN), h.PutUpdateBlog)
	appRouter.PATCH("/:id", middleware.RequireRole(user.ROLE_ADMIN), h.PathUpdateBlog)
	appRouter.POST("/:id/status", middleware.RequireRole(user.ROLE_ADMIN), h.UpdateBlogStatus)
	appRouter.DELETE("/:id", middleware.RequireRole(user.ROLE_ADMIN), h.DeleteBlog)

}

func (h *BlogApiHandler) QueryBlog(ctx *gin.Context) {
	// 获取用户请求
	req := blog.NewQueryBlogRequest()
	req.PageRequest = common.NewPageRequestFromGinCtx(ctx)
	req.KeyWords = ctx.Query("keywords")
	// 业务处理
	set, err := h.svc.QueryBlog(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(err, ctx)
		return
	}
	// 返回结果
	response.Success(set, ctx)
}
func (h *BlogApiHandler) DescribeBlog(ctx *gin.Context) {

	// 1.获取用户请求
	req := blog.NewDescribeBlogRequest(ctx.Param("id"))
	// 2.业务处理
	ins, err := h.svc.DescribeBlog(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(err, ctx)
		return
	}
	// 3.返回请求
	response.Success(ins, ctx)
}
func (h *BlogApiHandler) CreateBlog(ctx *gin.Context) {
	// 1.获取用户请求
	req := blog.NewCreateBlogRequest()
	if err := ctx.Bind(req); err != nil {
		response.Failed(exception.ErrValidateFailed(err.Error()), ctx)
		return
	}
	// 补充上下文中注入的中间数据
	if v, ok := ctx.Get(token.GIN_TOKEN_KEY_NAME); ok {
		req.CreateBy = v.(*token.Token).UserName
		if req.Author == "" {
			req.Author = req.CreateBy
		}
	}
	// 2.业务处理
	ins, err := h.svc.CreateBlog(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(err, ctx)
		return
	}
	// 3.返回请求
	response.Success(ins, ctx)
}
func (h *BlogApiHandler) PutUpdateBlog(ctx *gin.Context) {
	// 1.获取用户请求
	req := blog.NewUpdateBlogRequest(ctx.Param("id"))
	req.UpdateMode = common.UPDATE_MODE_PUT
	// body
	if err := ctx.Bind(req); err != nil {
		response.Failed(exception.ErrValidateFailed(err.Error()), ctx)
		return
	}
	// 2.业务处理
	ins, err := h.svc.UpdateBlog(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(err, ctx)
		return
	}
	// 3.返回请求
	response.Success(ins, ctx)
}
func (h *BlogApiHandler) PathUpdateBlog(ctx *gin.Context) {
	// 1.获取用户请求
	req := blog.NewUpdateBlogRequest(ctx.Param("id"))
	req.UpdateMode = common.UPDATE_MODE_PATCH
	// body
	if err := ctx.Bind(req); err != nil {
		response.Failed(exception.ErrValidateFailed(err.Error()), ctx)
		return
	}
	// 2.业务处理
	ins, err := h.svc.UpdateBlog(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(err, ctx)
		return
	}
	// 3.返回请求
	response.Success(ins, ctx)
}
func (h *BlogApiHandler) DeleteBlog(ctx *gin.Context) {
	// 1.获取用户请求
	req := blog.NewDeleteBlogRequest(ctx.Param("id"))
	// 2.业务处理
	ins, err := h.svc.DeleteBlog(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(err, ctx)
		return
	}
	// 3.返回请求
	response.Success(ins, ctx)

}
func (h *BlogApiHandler) UpdateBlogStatus(ctx *gin.Context) {
	// 1.获取用户请求
	req := blog.NewUpdateBlogStatusRequest(ctx.Param("id"))
	// body
	if err := ctx.Bind(req); err != nil {
		response.Failed(exception.ErrValidateFailed(err.Error()), ctx)
		return
	}
	// 2.业务处理
	ins, err := h.svc.UpdateBlogStatus(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(err, ctx)
		return
	}
	// 3.返回请求
	response.Success(ins, ctx)
}
