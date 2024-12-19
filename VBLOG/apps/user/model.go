package user

import (
	"encoding/json"
	"fmt"

	"gitee.com/VBLOG/common"
	"golang.org/x/crypto/bcrypt"
)

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Role:  ROLE_VISITOR,
		Label: map[string]string{},
	}
}
func NewUser(req *CreateUserRequest) *User {
	// hash密码

	return &User{
		Meta:              common.NewMeta(),
		CreateUserRequest: req,
	}
}

// 用户创建成功后返回一个User对象
// CreatedAt 为啥没用time.Time, int64(TimeStamp), 统一标准化, 避免时区你的程序产生影响
// 在需要对时间进行展示的时候，由前端根据具体展示那个时区的时间
type User struct {
	*common.Meta

	// 用户参数
	*CreateUserRequest
}

func (req *User) String() string {
	dj, _ := json.MarshalIndent(req, "", "	")
	return string(dj)
}

// 用户创建的参数
type CreateUserRequest struct {
	Username string `json:"username" validate:"required" gorm:"column:username"`
	Password string `json:"password" validate:"required" gorm:"column:password"`
	Role     Role   `json:"role" gorm:"column:role"`
	// https://gorm.io/docs/serializer.html
	// 用户标签 {"group": "a"} --json-> "{}"
	// 专门设计: label   id key value
	Label map[string]string `json:"label" gorm:"column:label;serializer:json"`
}

func (req *CreateUserRequest) Validate() error {
	if req.Username == "" {
		return fmt.Errorf("用户名必填")
	}
	return nil
}
func (req *CreateUserRequest) HashPassword() error {
	cryptoPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(cryptoPass)
	return nil
}

func (req *CreateUserRequest) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(password))
}

// 通用参数
type Meta struct {
	// 用户Id
	Id int `json:"id" gorm:"column:id"`
	// 创建时间, 时间戳 10位, 秒
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	// 更新时间, 时间戳 10位, 秒
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`
}

func NewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}
func (q *UserSet) String() string {
	dj, _ := json.MarshalIndent(q, "", " ")
	return string(dj)
}

// 一个对象的集合 UserCollection
type UserSet struct {
	// 总共有多个(分页,数据库里面总共)
	Total int64 `json:"total"`
	// 对象清单
	Items []*User `json:"items"`
}

func (req *User) TableName() string {
	return "users"
}
