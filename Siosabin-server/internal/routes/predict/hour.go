package predict

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hour(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
