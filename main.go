package main

import (
	"fmt"
	"net/http"

	"vincent.com/auth/services/jwt"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.POST("/api/auth/login", func(c *gin.Context) {
		type UID struct {
			ID string `json:"id"  binding:"required"`
		}
		var json UID
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// 此处默认生成uid过程
		token, err := jwt.New(json.ID)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})
	r.Run(":7000") // listen and serve on 0.0.0.0:8080
}
