# IOC(依赖倒置)：对象依赖管理

## 目前

```go
userServiceImpl := user.NewUserServiceImpl()
	// Token BO
	tokenServiceImpl := token.NewTokenServiceImpl(userServiceImpl)
	// 凡是路由有前缀是 /vblog/api/v1/tokens 到交给tokenApiHandler处理
	tokenApiHandler := api.NewTokenApiHandler(tokenServiceImpl)
```

在程序启动的时候(main)，收到传递依赖，main组装流程异常复杂，如果有2个模块(20 serviceImpl,20API对象)
main文件逻辑变得复杂

# IOC
IOC：含义使用逻辑或者角色发送了变化
被动模式：main , Developer ----》 Class: Developer 复杂依赖的管理；被动模式(class)，被动等待Develop在main组装时传递依赖
主动模式：ioc 就是一种主动模式，Class不再被动等待依赖传递过来，而是，自己主动到ioc容器里面去获取依赖，CLass() --> (Ioc获取依赖)模块开发，来获取自己依赖更加灵活集中

# IOC Container
系统内容所有业务对象(BO)的一个托儿所，这样业务对象的依赖才能自己去问IOC Container要，像Map{对象名称：对象地址}