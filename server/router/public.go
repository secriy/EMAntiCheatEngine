package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
)

// InitPublicRouter is public router
func InitPublicRouter(router *gin.RouterGroup) {
	BaseRouter := router.Group("")
	{
		BaseRouter.POST("user/login", v1.UserLogin) // User Login Api
		BaseRouter.POST("upload", v1.UploadLog)
	}
}
