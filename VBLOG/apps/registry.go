package apps

import (
	// 业务对象注册
	_ "gitee.com/VBLOG/apps/blog/api"
	_ "gitee.com/VBLOG/apps/blog/impl"
	_ "gitee.com/VBLOG/apps/token/api"
	_ "gitee.com/VBLOG/apps/token/impl"
	_ "gitee.com/VBLOG/apps/user/impl"
)
