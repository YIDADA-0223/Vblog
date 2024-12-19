package impl_test

import (
	"context"
	"crypto/md5"
	"fmt"
	"testing"

	"gitee.com/VBLOG/apps/user"
	"gitee.com/VBLOG/ioc"
	"gitee.com/VBLOG/test"
	"golang.org/x/crypto/bcrypt"
)

var (
	//声明被测试的对象
	serviceImpl user.Service
	ctx         = context.Background()
)

func init() {
	// 初始化单测环境
	test.DevelopmentSetup()
	// serviceImpl = impl.NewUserServiceImpl()
	serviceImpl = ioc.Controller.Get(user.AppName).(user.Service)
}

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "admin"
	req.Password = "123456"
	ins, err := serviceImpl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins.String())
}
func TestCreateVisitorUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "test"
	req.Password = "123456"
	req.Role = user.ROLE_ADMIN
	ins, err := serviceImpl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins.String())
}
func TestCreateAuthorUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "admin"
	req.Password = "123456"
	req.Role = user.ROLE_ADMIN
	ins, err := serviceImpl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins.String())
}
func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	req.Username = "admin"
	ins, err := serviceImpl.QueryUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
	// t.Log(ins.Items[0].CheckPassword("123456"))
}

func TestMd5(t *testing.T) {
	h := md5.New()
	h.Write([]byte("123456"))
	fmt.Printf("%x", h.Sum(nil))
}
func TestPasswordHash(t *testing.T) {
	password := "secret"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func TestUserCheckPassword(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "admin"
	req.Password = "123456"
	u := user.NewUser(req)

	u.HashPassword()
	t.Log(u.CheckPassword("123456"))
}
