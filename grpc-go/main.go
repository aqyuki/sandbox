package main

import (
	"context"
	"fmt"
	"gogrpc/pkg/proto"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterGreetingServiceServer(server, &GreetingServiceServerImpl{})
	reflection.Register(server)

	go func() {
		log.Println("start gRPC server on :8080")
		server.Serve(listener)
	}()

	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt)
	defer done()

	<-ctx.Done()
	log.Println("shutting down gRPC server...")
	server.GracefulStop()
}

type GreetingServiceServerImpl struct {
	// サービスンの後方互換性を保つために、自作のサービス構造体にはこの構造体を埋め込む
	proto.UnimplementedGreetingServiceServer
}

func (s *GreetingServiceServerImpl) Hello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
	}, nil
}
