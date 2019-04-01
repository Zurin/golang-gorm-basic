package master

import (
	"net/http"
	"new-platform-dashboard/db/repositories"
	"new-platform-dashboard/middleware"
	"new-platform-dashboard/service"
	"new-platform-dashboard/utils/builder"

	"github.com/gin-gonic/gin"
)

var cm *service.CacheManager

func UserRouter(r *gin.RouterGroup, permission middleware.Permission, cacheManager *service.CacheManager) {
	cm = cacheManager
	user := r.Group("/users")
	{
		user.POST("/save", permission.Set("PERMISSION_MASTER_USER_SAVE", UserSave))
		user.PUT("/update", permission.Set("PERMISSION_MASTER_USER_UPDATE", UserUpdate))
		user.GET("/find", permission.Set("PERMISSION_MASTER_USER_VIEW", UserFind))
		user.DELETE("delete/:id", permission.Set("PERMISSION_MASTER_USER_DELETE", UserDelete))
	}
}

func UserSave(c *gin.Context) {
	users, err := repositories.SaveUser(nil)
	if err != nil {
		c.JSON(http.StatusOK, builder.BaseResponse(false, err.Error(), nil))
	}

	cm.DeleteToken("#users")
	c.JSON(http.StatusOK, builder.BaseResponse(true, "ok", users))
}

func UserUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, builder.BaseResponse(true, "ok", nil))
}

func UserFind(c *gin.Context) {
	data := cm.Cacheable("#users", func() interface{} {
		users := repositories.FindUser(nil)
		return users
	})

	c.JSON(http.StatusOK, builder.BaseResponse(true, "ok", data))
}

func UserDelete(c *gin.Context) {
	c.JSON(http.StatusOK, builder.BaseResponse(true, "ok", nil))
}
