package ioc

// Controller是一个Controller，使用MapContainer的实现
var Controller Container = &MapContainer{
	name:   "controller",
	storge: make(map[string]Object),
}

// Api 所有的对外接口对象都放在这里
var Api Container = &MapContainer{
	name:   "api",
	storge: make(map[string]Object),
}
