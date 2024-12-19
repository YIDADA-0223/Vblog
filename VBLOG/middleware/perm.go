package middleware

import (
	"gitee.com/VBLOG/apps/token"
	"gitee.com/VBLOG/apps/user"
	"gitee.com/VBLOG/response"
	"github.com/gin-gonic/gin"
)

// 有参数中间件：Require("auditor")
// 通过一个函数返回一个中间件:gin.HandleFunc
// 这个中间件是加载在认证中间件之后的
func RequireRole(requiredRoles ...user.Role) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 判断当前用户的身份是不是与角色匹配
		// 补充上下文中注入的，中间数据
		if v, ok := ctx.Get(token.GIN_TOKEN_KEY_NAME); ok {
			for i := range requiredRoles {
				requiredRole := requiredRoles[i]
				if v.(*token.Token).Role == requiredRole {
					ctx.Next()
					return
				}
			}
		}
		response.Failed(token.ErrPermissionDenied, ctx)
		ctx.Abort()
	}
}
