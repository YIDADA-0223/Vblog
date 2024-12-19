package user

import (
	"context"

	"gitee.com/VBLOG/common"
)

const (
	// 业务包名称，用于托管这个业务包的业务对象，Service的具体实现
	AppName = "user"
)

type Service interface {
	//用户创建
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	//用户查询
	QueryUser(context.Context, *QueryUserRequest) (*UserSet, error)
}

func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{
		PageRequest: common.NewPageRequest(),
	}
}

type QueryUserRequest struct {
	Username string
	*common.PageRequest
}
