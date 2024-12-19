package token

import "context"

const (
	AppName = "token"
)

type Service interface {
	//令牌颁发
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	//令牌撤销（刷新令牌）
	RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error)
	//令牌校验,校验令牌合法性
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
}

func NewIssueTokenRequest(username, password string) *IssueTokenRequest {
	return &IssueTokenRequest{
		Username: username,
		Password: password,
		IsMember: false,
	}
}

type IssueTokenRequest struct {
	//用户密码
	Username string `json:"username"`
	Password string `json:"password"`
	//记住我
	IsMember bool `json:"is_member"`
}

func NewRevolkTokenRequest(at, rt string) *RevolkTokenRequest {
	return &RevolkTokenRequest{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

type RevolkTokenRequest struct {
	//撤销的令牌 （AccessToken，RefreshToken构成了一对username,password）
	AccessToken string
	//需要知道正确的刷新Token
	RefreshToken string
}

func NewValidateTokenRequest(at string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: at,
	}
}

type ValidateTokenRequest struct {
	//撤销的令牌 （AccessToken，RefreshToken构成了一对username,password）
	AccessToken string
}
