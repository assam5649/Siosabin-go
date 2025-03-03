package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Salinity(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
