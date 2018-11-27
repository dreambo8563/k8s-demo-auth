package main

import (
	"github.com/gin-gonic/gin"
	"vincent.com/auth/controllers"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.POST("/api/auth/login", controllers.JWTNewTokenHandler)
	r.Run(":7000") // listen and serve on 0.0.0.0:7000
}
