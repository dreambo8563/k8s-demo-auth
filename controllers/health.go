package controllers

import (
	"net/http"

	"go.uber.org/zap"

	"vincent.com/auth/services/logger"

	"github.com/gin-gonic/gin"
)

var log = logger.Logger

// HealthCheckHandler - health check handler
func HealthCheckHandler(c *gin.Context) {
	log.Info("HealthCheckHandler", zap.String("status", "ok"))
	// success
	c.JSON(http.StatusOK, "ok")
}
