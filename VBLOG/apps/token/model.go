package token

import (
	"encoding/json"
	"time"

	"gitee.com/VBLOG/apps/user"
	"github.com/rs/xid"
)

// 构造函数的业务逻辑
// 颁发给用户的Token，给user对象颁发
func NewToken(u *user.User) *Token {
	return &Token{
		UserId:   u.Id,
		UserName: u.Username,
		//使用随机UUID
		AccessToken: xid.New().String(),
		//过期Token
		AccessTokenExpiredAt:  3600,
		RefreshToken:          xid.New().String(),
		RefreshTokenExpiredAt: 3600 * 4,
		CreatedAt:             time.Now().Unix(),
		Role:                  u.Role,
	}
}

func DefaultToken() *Token {
	return &Token{}
}

// 这个也相当于定义接口都需要返回什么样的参数
type Token struct {
	// 该Token是颁发
	UserId int `json:"user_id" gorm:"column:user_id"`
	// 人的名称， user_name
	UserName string `json:"username" gorm:"column:username"`
	// 办法给用户的访问令牌(用户需要携带Token来访问接口)
	AccessToken string `json:"access_token" gorm:"column:access_token"`
	// 过期时间(2h), 单位是秒
	AccessTokenExpiredAt int `json:"access_token_expired_at" gorm:"column:access_token_expired_at"`
	// 刷新Token
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token"`
	// 刷新Token过期时间(7d)
	RefreshTokenExpiredAt int `json:"refresh_token_expired_at" gorm:"column:refresh_token_expired_at"`
	// 创建时间
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	// 更新实现
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	// 额外补充信息, gorm忽略处理
	Role user.Role `json:"role" gorm:"-"`
}

func (t *Token) IssueTime() time.Time {
	return time.Unix(t.CreatedAt, 0)
}
func (t *Token) AccessTokenDuration() time.Duration {
	return time.Duration(t.AccessTokenExpiredAt * int(time.Second))
}

func (t *Token) RefreshTokenDuration() time.Duration {
	return time.Duration(time.Duration(t.RefreshTokenExpiredAt * int(time.Second)))
}

// 颁发Token
func (t *Token) AccessTokenIsExpired() error {
	//过期时间: 颁发时间　＋　过期时长
	expiredTime := t.IssueTime().Add(t.AccessTokenDuration())
	//当前时间减去过期时间
	if time.Since(expiredTime).Seconds() > 0 {
		return ErrAccessTokenExpired
	}
	return nil
}

// 过期Token
func (t *Token) RefreshTokenIsExpired() error {
	//过期时间: 颁发时间　＋　过期时长
	expiredTime := t.IssueTime().Add(t.RefreshTokenDuration())
	//当前时间减去过期时间
	if time.Since(expiredTime).Seconds() > 0 {
		return RefreshTokenExpired
	}
	return nil
}
func (t *Token) String() string {
	dj, _ := json.MarshalIndent(t, "", " ")
	return string(dj)
}
