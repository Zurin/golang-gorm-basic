package public

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
}
