package routes

import (
	"github.com/assam5649/Siosabin-go/Siosabin-server/internal/routes/auth"
	"github.com/assam5649/Siosabin-go/Siosabin-server/internal/routes/predict"
	"github.com/assam5649/Siosabin-go/Siosabin-server/internal/routes/user"

	"github.com/gin-gonic/gin"
)

func NewRouter() {

	r := gin.Default()

	r.GET("/", auth.Ping)
	r.GET("/", user.Ping)
	r.GET("/", predict.Ping)

	r.Run(":8080")
}
