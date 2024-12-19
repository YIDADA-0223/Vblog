package token

import (
	"net/http"

	"gitee.com/VBLOG/exception"
)

var (
	ErrUnauthorized       = exception.NewApiExcepiton(50000, "请登录").WithHttpCode(http.StatusUnauthorized)
	ErrAuthFailed         = exception.NewApiExcepiton(50001, "用户名或者密码不正确").WithHttpCode(http.StatusUnauthorized)
	ErrAccessTokenExpired = exception.NewApiExcepiton(50002, "AccessToken过期")
	RefreshTokenExpired   = exception.NewApiExcepiton(50003, "RefreshToken过期")
	ErrPermissionDenied   = exception.NewApiExcepiton(50004, "该角色无法访问当前接口").WithHttpCode(http.StatusForbidden)
)
