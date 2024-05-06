package grpc

import (
	"context"
	"log"
	"net"

	proto "github.com/dev-crusader404/go-test-project/grpc/protogen"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	proto.UnimplementedMovieInterfaceServer
}

func InitServer(addr string) {

	lis, err := net.Listen("tcp/ip", addr)
	if err != nil {
		log.Fatalf("error starting up server: %v", err)
		return
	}
	opts := []grpc.ServerOption{}
	sv := grpc.NewServer(opts...)
	grpcService := &GrpcServer{}
	proto.RegisterMovieInterfaceServer(sv, grpcService)
	err = sv.Serve(lis)
	if err != nil {
		log.Fatalf("error serving: %v", err)
		return
	}
	log.Printf("grpc server running on port %s", addr)
}

func (s *GrpcServer) SearchMovie(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	return &proto.Response{}, nil
}
