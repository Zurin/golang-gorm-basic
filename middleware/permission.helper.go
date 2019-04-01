package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new-platform-dashboard/utils/builder"
)

type Permission interface {
	Set(name string, callback func(*gin.Context)) gin.HandlerFunc
}

type Permit struct {
}

func (p Permit) Set(name string, callback func(*gin.Context)) gin.HandlerFunc {
	if name == "PERMISSION_MASTER_USER_VIEW" || name == "PERMISSION_MASTER_USER_SAVE" {
		return callback
	}

	return func(c *gin.Context) {
		c.JSON(http.StatusForbidden, builder.BaseResponse(false, http.StatusText(http.StatusForbidden), nil))
	}
}
