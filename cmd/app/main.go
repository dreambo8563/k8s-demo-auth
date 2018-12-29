package main

import (
	"github.com/gin-gonic/gin"
	"vincent.com/auth/internal/controllers"
	"vincent.com/auth/internal/pkg/tracing"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	tracing.Init("todo-auth-service")
	defer tracing.Closer.Close()

	r := gin.Default()
	go controllers.InitRPCServer(tracing.Tracer)
	r.POST("/api/auth/login", controllers.JWTNewTokenHandler)
	r.GET("/healthz", controllers.HealthCheckHandler)
	r.Run(":7000") // listen and serve on 0.0.0.0:7000
}
