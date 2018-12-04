package controllers

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"vincent.com/auth/rpc/helloworld"
	"vincent.com/auth/services/logger"
	"vincent.com/auth/services/tracing"

	"github.com/gin-gonic/gin"
)

const (
	port = ":50051"
)

var log = logger.Logger

// server is used to implement helloworld.GreeterServer.
type server struct{}

// HealthCheckHandler - health check handler
func HealthCheckHandler(c *gin.Context) {
	span := tracing.Tracer.StartSpan("HealthCheckHandler")
	defer span.Finish()
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
	helloworld.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Sugar().Fatalf("failed to serve: %v", err)
	}
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "SayHello")
	defer span.Finish()
	return &helloworld.HelloReply{Message: "Hello " + in.Name}, nil
}
