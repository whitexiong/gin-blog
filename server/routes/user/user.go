package user

import (
	controller "blog/app/controllers/user"

	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	// e.GET("/user/create", CreateHandler)
	// e.POST("/user/login", LoginHandler)
	e.POST("/login", controller.Login) // 登录
	// e.POST("/socialLogin", SocialLoginHandler)
	e.POST("/userInfo", controller.UserInfo)     // 获取用户信息
	e.POST("/logout", controller.Logout)         // 退出登录
	e.POST("/register", controller.RegisterUser) // 注册
}
