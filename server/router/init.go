package router

import (
	"github.com/gin-gonic/gin"
	"server/conf"
	"server/middleware"
)

// NewRouter Create New Router
func NewRouter() *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(middleware.Session(conf.ReadConfig().SessionSecret))
	router.Use(middleware.Cors())
	router.Use(middleware.CurrentUser())

	// Router Group
	r := router.Group("api/v1")

	// Public Router
	InitPublicRouter(r)

	// User Router
	InitUserRouter(r)

	// Log Router
	InitLogRouter(r)

	return router
}
