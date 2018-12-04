package controllers

import (
	"context"
	"net"
	"net/http"

	"vincent.com/auth/services/tracing"

	"vincent.com/auth/rpc/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"vincent.com/auth/services/logger"

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
func InitRPCServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Sugar().Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Sugar().Fatalf("failed to serve: %v", err)
	}
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.Name}, nil
}
