package controllers

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"vincent.com/auth/services/jwt"
	"vincent.com/auth/services/tracing"
)

// JWTNewTokenHandler - new token handler
func JWTNewTokenHandler(c *gin.Context) {
	tracer := tracing.Tracer
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	span := tracer.StartSpan("JWTNewTokenHandler", ext.RPCServerOption(spanCtx))
	fmt.Println(c.Request.Header)
	defer span.Finish()
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
	span.SetTag("UID", reqParams.ID)
	log.Sugar().Infow("get reqParams", "params", reqParams)
	span.LogKV("event", "verified params", "params", reqParams)
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	token, err := jwt.New(ctx, reqParams.ID)
	if err != nil {
		// jwt err
		span.LogKV("event", "jwt err", "err", err)
		log.Error("generate token err", zap.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	log.Info("receive token", zap.String("token", token))
	span.LogKV("event", "jwt success", "token", token)
	// success
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
