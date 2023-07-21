package error

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(c *gin.Context, err error) bool {
	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return true
	}
	return false
}
