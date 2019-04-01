package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"new-platform-dashboard/api/master"
	"new-platform-dashboard/api/public"
	"new-platform-dashboard/middleware"
	"new-platform-dashboard/service"
)

//registering router handler
func Init(router *gin.Engine, cache *service.CacheManager) {
	permission := middleware.Permit{}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.NoRoute(public.NotFound)

	v1 := router.Group("/api/v1")
	{
		master.UserRouter(v1, permission, cache)
		public.LoginRouter(v1, permission)
	}
}
