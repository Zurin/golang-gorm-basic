package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new-platform-dashboard/middleware"
	"new-platform-dashboard/utils/builder"
)

func LoginRouter(r *gin.RouterGroup, permission middleware.Permission) {
	r.GET("/login", login)
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, builder.BaseResponse(true, "ok", nil))
}
