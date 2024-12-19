package api

import (
	"gitee.com/VBLOG/apps/token"
	"gitee.com/VBLOG/conf"
	"gitee.com/VBLOG/ioc"
	"gitee.com/VBLOG/response"
	"github.com/gin-gonic/gin"
)

func NewTokenApiHandler() *TokenApiHandler {
	return &TokenApiHandler{
		token: ioc.Controller.Get(token.AppName).(token.Service),
	}
}

// 注册api
func init() {
	ioc.Api.Registry(token.AppName, &TokenApiHandler{})
}

// 核心处理APi请求
type TokenApiHandler struct {
	// 依赖TokenServiceImpl
	// impl.TokenServiceImpl 不推荐这样使用
	//依赖接口
	token token.Service
}

// 自己决定初始化逻辑
func (h *TokenApiHandler) Init() error {
	h.token = ioc.Controller.Get(token.AppName).(token.Service)
	// 已经前置在config中/vblog/v1/
	subRouter := conf.C().Application.GinRootRouter().Group("tokens")
	h.Registry(subRouter)
	return nil
}

// 把自己的路由信息注册给Gin Root Router
// 每个业务模块，有每个业务模块的子路由
func (h *TokenApiHandler) Registry(appRouter gin.IRouter) {
	// r := gin.Default()
	// //RouterGroup
	// r.Group("api").Group("v1")
	// POST /api/v1/tokens/ ---> Login
	appRouter.POST("/", h.Login)
	appRouter.DELETE("/", h.Logout)
}

// 颁发令牌
// Gin Hander Http Response Http Request
func (h *TokenApiHandler) Login(c *gin.Context) {
	//1.获取HTTP请求
	req := token.NewIssueTokenRequest("", "")
	if err := c.BindJSON(req); err != nil {
		response.Failed(err, c)
		return
	}
	//2.业务处理
	tk, err := h.token.IssueToken(c.Request.Context(), req)
	if err != nil {
		//不是用的我们的自定义异常
		response.Failed(err, c)
		return
	}
	c.SetCookie(
		token.COOKIE_TOKEY_KEY,
		tk.AccessToken,
		tk.RefreshTokenExpiredAt,
		"/", conf.C().Application.Domain,
		false,
		true)
	response.Success(tk, c)
	// c.JSON(http.StatusOK, tk)
	//3.返回结果

}

func (h *TokenApiHandler) Logout(c *gin.Context) {
	// 1.获取HTTP请求 IssueTokenRequest
	// DELETE方法一般情况不带Body
	// 敏感信息放Header或者Body
	ak, err := c.Cookie(token.COOKIE_TOKEY_KEY)
	if err != nil {
		response.Failed(err, c)
		return
	}
	rt := c.GetHeader(token.REFRESH_HEADER_KEY)
	req := token.NewRevolkTokenRequest(ak, rt)
	tk, err := h.token.RevolkToken(c.Request.Context(), req)
	if err != nil {
		response.Failed(err, c)
		return
	}
	response.Success(tk, c)
}
