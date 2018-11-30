package controllers

import (
	"net/http"

	"vincent.com/auth/services/logger"

	"github.com/gin-gonic/gin"
)

var log = logger.Logger

// HealthCheckHandler - health check handler
func HealthCheckHandler(c *gin.Context) {
	// success
	c.JSON(http.StatusOK, "ok")
}
