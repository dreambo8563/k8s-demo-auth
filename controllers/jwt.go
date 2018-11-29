package controllers

import (
	"net/http"

	"go.uber.org/zap"

	"vincent.com/auth/services/logger"

	"vincent.com/auth/services/jwt"

	"github.com/gin-gonic/gin"
)

var log = logger.Logger

// JWTNewTokenHandler - new token handler
func JWTNewTokenHandler(c *gin.Context) {
	var reqParams struct {
		ID string `json:"id"  binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqParams); err != nil {
		// params err
		log.Sugar().Warn("JWTNewTokenHandler params err", reqParams)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	log.Sugar().Infow("get reqParams", "params", reqParams)
	token, err := jwt.New(reqParams.ID)
	if err != nil {
		// jwt err
		log.Error("generate token err", zap.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	log.Info("receive token", zap.String("token", token))

	// success
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
