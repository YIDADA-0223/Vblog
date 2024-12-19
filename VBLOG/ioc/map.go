package ioc

import (
	"fmt"
)

type MapContainer struct {
	name   string
	storge map[string]Object
}

// 注册对象
func (c *MapContainer) Registry(name string, obj Object) {
	c.storge[name] = obj
}

// 获取对象
func (c *MapContainer) Get(name string) any {
	return c.storge[name]
}

// 调用所有被托管对象的init方法，对像进行初始化
func (c *MapContainer) Init() error {
	for k, v := range c.storge {
		if err := v.Init(); err != nil {
			return fmt.Errorf("%s init error,%s", k, err)
		}
		fmt.Printf("[%s] %s init successs \n", c.name, k)
	}
	return nil
}

// // 路由注册,转一个Root进来，把对象上所有
// func (c *MapContainer) RegistryGinRoute(root gin.IRouter) {
// 	for k, v := range c.storge {

// 	}
// }
