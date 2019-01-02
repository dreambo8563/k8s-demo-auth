package controllers

import (
	"context"
	"net"

	"vincent.com/auth/internal/domain/usecase"

	"vincent.com/auth/internal/adapter/service"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"

	opentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"vincent.com/auth/internal/adapter/http/rpc/auth"
)

// JWTNewTokenHandler - new token handler
// func JWTNewTokenHandler(c *gin.Context) {
// 	tracer := tracing.Tracer
// 	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
// 	span := tracer.StartSpan("JWTNewTokenHandler", ext.RPCServerOption(spanCtx))
// 	fmt.Println(c.Request.Header)
// 	defer span.Finish()
// 	var reqParams struct {
// 		ID string `json:"id"  binding:"required"`
// 	}

// 	if err := c.ShouldBindJSON(&reqParams); err != nil {
// 		// params err
// 		log.Sugar().Warn("JWTNewTokenHandler params err", reqParams)
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": err.Error(),
// 		})
// 		return
// 	}
// 	span.SetTag("UID", reqParams.ID)
// 	log.Sugar().Infow("get reqParams", "params", reqParams)
// 	span.LogKV("event", "verified params", "params", reqParams)
// 	ctx := opentracing.ContextWithSpan(context.Background(), span)
// 	token, err := jwt.New(ctx, reqParams.ID)
// 	if err != nil {
// 		// jwt err
// 		span.LogKV("event", "jwt err", "err", err)
// 		log.Error("generate token err", zap.String("error", err.Error()))
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"msg": err.Error(),
// 		})
// 		return
// 	}
// 	log.Info("receive token", zap.String("token", token))
// 	span.LogKV("event", "jwt success", "token", token)
// 	// success
// 	c.JSON(http.StatusOK, gin.H{
// 		"token": token,
// 	})
// }

// InitRPCServer -
func InitRPCServer(tracer opentracing.Tracer) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Sugar().Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(
		otgrpc.OpenTracingServerInterceptor(tracer)),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer)))
	auth.RegisterAuthServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Sugar().Fatalf("failed to serve: %v", err)
	}
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetToken(ctx context.Context, in *auth.GetTokenRequest) (*auth.GetTokenReply, error) {
	span, childCtx := opentracing.StartSpanFromContext(ctx, "SayHello")
	defer span.Finish()
	span.SetTag("UID", in.Uid)
	authcase := service.InitializeAuthCase()

	token, err := authcase.NewToken(childCtx, &usecase.User{
		ID: in.Uid,
	})
	if err != nil {
		// jwt err
		span.LogKV("event", "jwt err", "err", err)
		log.Error("generate token err", log.String("error", err.Error()))

		return nil, err
	}
	span.LogKV("event", "jwt success", "token", token)
	return &auth.GetTokenReply{Token: "Hello " + token}, nil
}
