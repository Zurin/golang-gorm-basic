package middleware

import (
	"github.com/gin-gonic/gin"
	"new-platform-dashboard/service"
)

func Auth(cache *service.CacheManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
