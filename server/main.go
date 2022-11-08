package main

import (
	"blog/app/user"
	"blog/routes"
	"fmt"
)

func main() {
	// 加载多个APP的路由配置
	routes.Include(user.Routers)
	// 初始化路由
	r := routes.Init()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
