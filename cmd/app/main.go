package main

import (
	"github.com/gin-gonic/gin"
	"vincent.com/auth/internal/controllers"
	"vincent.com/auth/internal/pkg/tracing"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	traceClient := tracing.NewTraceClient()
	defer traceClient.Closer.Close()

	r := gin.Default()
	go controllers.InitRPCServer(traceClient.GetTracer())
	r.GET("/healthz", controllers.HealthCheckHandler)
	r.Run(":7000") // listen and serve on 0.0.0.0:7000
}
