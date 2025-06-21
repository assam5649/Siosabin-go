package routes

import (
	"Siosabin-go/Siosabin-server/internal/routes/auth"
	"Siosabin-go/Siosabin-server/internal/routes/predict"
	"Siosabin-go/Siosabin-server/internal/routes/user"

	"github.com/gin-gonic/gin"
)

func NewRouter() {

	r := gin.Default()

	r.GET("/auth/", auth.Ping)
	r.POST("/auth/login", auth.Login)
	r.POST("/auth/register", auth.Register)

	r.GET("/main/", user.Ping)
	r.POST("/main/data", user.Data)
	r.GET("/main/users", user.Users)
	r.GET("/main/location", user.Location)
	r.GET("/main/salinity", user.Salinity)

	r.GET("/predict/", predict.Ping)
	r.GET("/predict/hour", predict.Hour)
	r.GET("/predict/day", predict.Day)

	r.Run(":8080")
}
