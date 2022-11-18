package main

import (
	"blog/common"
	"blog/initialize"
)

func main() {

	defer common.DB.Close()
	routers := initialize.InitRouters()
	_ = routers.Run(":" + common.CONFIG.System.Port)
}

func init() {
	//初始华配置
	initialize.InitConf()
	// 初始化日志
	initialize.InitLog()
	//初始化redis
	// initialize.InitCache()
	//初始化数据库
	initialize.InitDb()
	// 初始化Casbin
	initialize.InitCasbin()
}
