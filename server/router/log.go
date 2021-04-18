package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

// InitLogRouter is log private router
func InitLogRouter(router *gin.RouterGroup) {
	LogRouter := router.Group("").Use(middleware.AuthRequired())
	{
		LogRouter.GET("list/anti", v1.ListAntiLog)
		LogRouter.GET("list/hook", v1.ListHookLog)
		LogRouter.GET("list/anti/exp", v1.ListExpAntiLog)
		LogRouter.GET("list/hook/exp", v1.ListExpHookLog)
		LogRouter.GET("count/anti", v1.CountAntiLog)
		LogRouter.GET("count/hook", v1.CountHookLog)
	}
}
