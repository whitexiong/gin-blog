package routers

import (
	"blog/middleware"

	v1 "blog/app/admin"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	// 用户登录
	Router.POST("/login", v1.Login)
	Router.POST("/logout", middleware.JWTAuth(), v1.LogOut)
}
