package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Data(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
