package token_test

import (
	"testing"
	"time"

	"gitee.com/VBLOG/apps/token"
	"gitee.com/VBLOG/apps/user"
)

func TestTokenJson(t *testing.T) {
	tk := token.Token{
		UserId:   1,
		UserName: "admin",
	}
	t.Log(tk)
}
func TestTokenExired(t *testing.T) {
	now := time.Now().Unix()
	tk := token.Token{
		UserId:               1,
		Role:                 user.ROLE_ADMIN,
		AccessTokenExpiredAt: 1,
		CreatedAt:            now,
	}
	t.Log(tk.AccessTokenIsExpired())
}
