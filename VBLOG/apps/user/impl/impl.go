package impl

import (
	"context"

	"gitee.com/VBLOG/apps/user"
	"gitee.com/VBLOG/common"
	"gitee.com/VBLOG/conf"
	"gitee.com/VBLOG/ioc"
	"gorm.io/gorm"
)

// func NewUserServiceImpl() *UserServiceImpl {
// 	return &UserServiceImpl{
// 		db: conf.C().MySQL.GetDB(),
// 	}
// }

// import ---> init方法来注册包里面的核心对象
func init() {
	ioc.Controller.Registry(user.AppName, &UserServiceImpl{})
}

type UserServiceImpl struct {
	db *gorm.DB
}

// 对象属性初始化
func (i *UserServiceImpl) Init() error {
	i.db = conf.C().MySQL.GetDB()
	return nil
}

func (i *UserServiceImpl) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.User, error) {
	if err := common.Validate(in); err != nil {
		return nil, err
	}
	if err := in.HashPassword(); err != nil {
		return nil, err
	}
	ins := user.NewUser(in)

	if err := i.db.WithContext(ctx).Save(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *UserServiceImpl) QueryUser(ctx context.Context, in *user.QueryUserRequest) (*user.UserSet, error) {
	set := user.NewUserSet()

	// 构造一个查询语句, TableName() select
	// WithContext
	query := i.db.Model(&user.User{}).WithContext(ctx)

	// Where where username = ?
	// SELECT * FROM `users` WHERE username = 'admin' LIMIT 10
	if in.Username != "" {
		// 注意: 返回一个新的对象, 并没有直接修改对象
		// 新生产的query 语句才有 query
		query = query.Where("username = ?", in.Username)
	}
	// ...

	// 怎么查询Total, 需要把过滤条件: username ,key
	// 查询Total时能不能把分页参数带上
	// select COUNT(*) from xxx limit 10
	// select COUNT(*) from xxx
	// 不能携带分页参数
	if err := query.Count(&set.Total).Error; err != nil {
		return nil, err
	}

	// LIMIT ?,?
	// SELECT * FROM `users` LIMIT 10
	if err := query.
		Offset(in.Offset()).
		Limit(in.PageSize).
		Find(&set.Items).Error; err != nil {
		return nil, err
	}

	return set, nil
}
