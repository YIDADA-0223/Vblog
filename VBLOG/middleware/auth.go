package middleware

import (
	"gitee.com/VBLOG/apps/token"
	"gitee.com/VBLOG/ioc"
	"gitee.com/VBLOG/response"
	"github.com/gin-gonic/gin"
)

// Gin Web中间件，我们需要在中间件注入到请求的链路当中，然后由Gin框架来调用

func Auth(ctx *gin.Context) {
	accessToken, err := ctx.Cookie(token.COOKIE_TOKEY_KEY)
	if err != nil {
		response.Failed(token.ErrUnauthorized.WithMessage(err.Error()), ctx)
		ctx.Abort()
	}
	tk, err := ioc.Controller.Get(token.AppName).(token.Service).ValidateToken(ctx.Request.Context(), token.NewValidateTokenRequest(accessToken))
	if err != nil {
		// 响应报错信息
		response.Failed(token.ErrAuthFailed.WithMessage(err.Error()), ctx)
		ctx.Abort()
	} else {
		// 鉴权成功，请求继续往后面进行
		// Gin 采用一个map对象来维护中间传递的数据
		ctx.Set(token.GIN_TOKEN_KEY_NAME, tk)
		ctx.Next()
	}
}
