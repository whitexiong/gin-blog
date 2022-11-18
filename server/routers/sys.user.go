package routers

import (
	v1 "blog/app/admin"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserGroup := Router.Group("admin").Use(middleware.JWTAuth())
	{
		UserGroup.GET("/user/list", v1.GetUserList)
		UserGroup.POST("/user/add", v1.CreateUser)
		UserGroup.PUT("/user/edit", v1.UpdateUser)
		UserGroup.DELETE("/user/del", v1.DelUser)
	}
}
