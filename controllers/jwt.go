package controllers

import (
	"net/http"

	"vincent.com/auth/services/jwt"

	"github.com/gin-gonic/gin"
)

// JWTNewTokenHandler - new token handler
func JWTNewTokenHandler(c *gin.Context) {
	var reqParams struct {
		ID string `json:"id"  binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqParams); err != nil {
		// params err
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	token, err := jwt.New(reqParams.ID)
	if err != nil {
		// jwt err
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// success
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
