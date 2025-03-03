package predict

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Day(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
