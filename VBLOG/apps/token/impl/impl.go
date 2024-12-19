package impl

import (
	"context"
	"fmt"

	"gitee.com/VBLOG/apps/token"
	"gitee.com/VBLOG/apps/user"
	"gitee.com/VBLOG/conf"
	"gitee.com/VBLOG/exception"
	"gitee.com/VBLOG/ioc"
	"gorm.io/gorm"
)

//	func NewTokenServiceImpl() *TokenServiceImpl {
//		//获取一个全新的mysql连接池对象
//		//程序启动的时候一定要加载配置
//		return &TokenServiceImpl{
//			db:   conf.C().MySQL.GetDB(),
//			user: ioc.Controller.Get(user.AppName).(user.Service),
//		}
//	}
func init() {
	ioc.Controller.Registry(token.AppName, &TokenServiceImpl{})
}

// Token接口的实现，存放DB
type TokenServiceImpl struct {
	db *gorm.DB
	//依赖用户服务
	user user.Service
}

func (i *TokenServiceImpl) Init() error {
	i.db = conf.C().MySQL.GetDB()
	// 获取对象Controller.Get
	// 断言对象实现了user.Service
	i.user = ioc.Controller.Get(user.AppName).(user.Service)
	return nil
}

// 令牌颁发
func (i *TokenServiceImpl) IssueToken(ctx context.Context, in *token.IssueTokenRequest) (*token.Token, error) {
	//1.查询用户对象
	queryUser := user.NewQueryUserRequest()
	queryUser.Username = in.Username
	us, err := i.user.QueryUser(ctx, queryUser)
	if err != nil {
		return nil, err
	}
	if len(us.Items) == 0 {
		//安全规范，避免爆破
		return nil, token.ErrAuthFailed
	}

	//2.比对用户密码
	u := us.Items[0]
	if err := us.Items[0].CheckPassword(in.Password); err != nil {
		return nil, token.ErrAuthFailed
	}

	//3.上面密码比对成功，颁发令牌Token,UUID,自己随机生成一个多少位字符串
	tk := token.NewToken(u)
	//4.把令牌存储在数据库里面
	//INSERT INTO `tokens` (`user_id`,`username`,`access_token`,`access_token_expired_at`,`refresh_token`,`refresh_token_expired_at`,`created_at`,`updated_at`) VALUES (1,'admin','cpje9nols52l1e13vn0g',3600,'cpje9nols52l1e13vn10',14400,1718019295,1718019295)
	if err := i.db.WithContext(ctx).Create(tk).Error; err != nil {
		return nil, exception.ErrServerInternal("保存报错,%s", err)
	}
	return tk, nil
}

// 令牌撤销
func (i *TokenServiceImpl) RevolkToken(ctx context.Context, in *token.RevolkTokenRequest) (*token.Token, error) {
	//ErrServerInternal没有使用
	//后面可以用中间件来统一处理：非 ApiExcepiton
	//直接删除数据里面存储的Token
	tk := token.DefaultToken()
	//DELETE FROM `tokens` WHERE access_token = 'cpje9nols52l1e13vn0g'
	err := i.db.WithContext(ctx).Where("access_token = ?", in.AccessToken).First(tk).Error
	if err == gorm.ErrRecordNotFound {
		return nil, exception.ErrNotFound("Token未找到")
	}
	if tk.RefreshToken != in.RefreshToken {
		return nil, fmt.Errorf("RefreshToken不正确")
	}
	// DELETE FROM `tokens` WHERE access_token = 'cpijm58ls52k3i39bukg'
	err = i.db.WithContext(ctx).Where("access_token = ?", in.AccessToken).Delete(token.Token{}).Error
	if err != nil {
		return nil, err
	}
	return tk, nil
}

// 令牌校验,校验令牌合法性
func (i *TokenServiceImpl) ValidateToken(ctx context.Context, in *token.ValidateTokenRequest) (*token.Token, error) {
	//1.查询出Token，where accessToken来查询
	tk := token.DefaultToken()
	err := i.db.WithContext(ctx).Where("access_token = ?", in.AccessToken).First(tk).Error
	if err == gorm.ErrRecordNotFound {
		return nil, exception.ErrNotFound("Token未找到")
	}
	if err != nil {
		return nil, exception.ErrServerInternal("查询报错,%v", err)
	}
	//2.判断Token是否过期，（1）判断RefreshToken有没有过期，（2）AccessToken有没有过期
	if err := tk.RefreshTokenIsExpired(); err != nil {
		return nil, err
	}
	if err := tk.AccessTokenIsExpired(); err != nil {
		return nil, err
	}
	//3.Token合法
	//4.补充用户角色信息
	queryUserReq := user.NewQueryUserRequest()
	queryUserReq.Username = tk.UserName
	us, err := i.user.QueryUser(ctx, queryUserReq)
	if err != nil {
		return nil, err
	}
	if len(us.Items) == 0 {
		return nil, fmt.Errorf("token user not found")
	}
	tk.Role = us.Items[0].Role
	return tk, nil
}
