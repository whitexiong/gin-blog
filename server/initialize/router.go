package initialize

import (
	"blog/common"
	"blog/middleware"
	"blog/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouters() (Router *gin.Engine) {
	Router = gin.Default()

	Router.Use(middleware.Cors())
	common.LOG.Info("use middleware cors logger recovery")

	Router.Static("/dist", "../web/dist")
	//Router.LoadHTMLGlob("templates/*")
	Router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/dist")
	})

	ApiV1Group := Router.Group("/api")

	routers.InitUserRouter(ApiV1Group)
	common.LOG.Info("routers register success")

	return
}
