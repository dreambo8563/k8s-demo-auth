package controllers

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"vincent.com/auth/internal/pkg/jwt"
	"vincent.com/auth/internal/pkg/logger"
	"vincent.com/auth/internal/rpc/auth"

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
	token, err := jwt.New(childCtx, in.Uid)
	if err != nil {
		// jwt err
		span.LogKV("event", "jwt err", "err", err)
		log.Error("generate token err", zap.String("error", err.Error()))

		return nil, err
	}
	span.LogKV("event", "jwt success", "token", token)
	return &auth.GetTokenReply{Token: "Hello " + token}, nil
}
