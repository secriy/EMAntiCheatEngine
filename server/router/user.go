package router

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

// InitUserRouter is user private router
func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("user").Use(middleware.AuthRequired())
	{
		UserRouter.GET("me", v1.UserMe)
		UserRouter.DELETE("logout", v1.UserLogout)
		UserRouter.PUT("pass", v1.ChangePass)
		UserRouter.GET("all", v1.UsersList)
		UserRouter.POST("add", v1.UserAdd)
		UserRouter.DELETE("del/:id", v1.UserDelete)
		UserRouter.PUT("update/:id", v1.UserUpdate)
	}
}
