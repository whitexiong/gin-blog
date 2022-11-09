package main

import (
	"blog/app/dao"
	"blog/app/models"
	"blog/app/setting"
	"blog/routes"
	"blog/routes/user"
	"fmt"
)

func main() {

	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	// 模型绑定
	dao.DB.AutoMigrate(&models.User{})

	// 加载多个APP的路由配置
	routes.Include(user.Routers)
	// 初始化路由
	r := routes.Init()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
