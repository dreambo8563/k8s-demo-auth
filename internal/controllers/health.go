package controllers

import (
	"net/http"

	"vincent.com/auth/internal/pkg/logger"

	"github.com/gin-gonic/gin"
)

const (
	port = ":50051"
)

var log = logger.Logger()

// server is used to implement helloworld.GreeterServer.
type server struct{}

// HealthCheckHandler - health check handler
func HealthCheckHandler(c *gin.Context) {
	// span := tracing.Tracer.StartSpan("HealthCheckHandler")
	// defer span.Finish()
	// success
	c.JSON(http.StatusOK, "ok")
}
