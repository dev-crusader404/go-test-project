package server

import (
	"context"
	"log"
	"net"

	"github.com/dev-crusader404/go-test-project/client"
	proto "github.com/dev-crusader404/go-test-project/grpc/protogen"
	mv "github.com/dev-crusader404/go-test-project/internal"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	proto.UnimplementedMovieInterfaceServer
	fetcher mv.MovieFetcher
}

func NewGrpcServer(fetcher mv.MovieFetcher) *GrpcServer {
	return &GrpcServer{
		fetcher: fetcher,
	}
}

func InitServer(addr string) {

	lis, err := net.Listen("tcp/ip", addr)
	if err != nil {
		log.Fatalf("error starting up server: %v", err)
		return
	}
	opts := []grpc.ServerOption{}
	sv := grpc.NewServer(opts...)
	m := &mv.Movie{Client: client.GetClient()}
	grpcService := NewGrpcServer(m)
	proto.RegisterMovieInterfaceServer(sv, grpcService)
	err = sv.Serve(lis)
	if err != nil {
		log.Fatalf("error serving: %v", err)
		return
	}
	log.Printf("grpc server running on port %s", addr)
}

func (s *GrpcServer) SearchMovie(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	log.Printf("\n Received SearchMovie request: %+v", req)
	id, err := s.fetcher.GetMovie(ctx, req.Title, req.Year)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return nil, err
	}

	result, err := s.fetcher.GetDetails(ctx, id)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return nil, err
	}
	response := proto.Response{
		MovieTitle:   result.MovieTitle,
		Year:         result.Year,
		Description:  result.Overview,
		Rating:       result.Rating,
		Genre:        result.Genre,
		ReleasedDate: result.ReleasedDate,
		GrossIncome:  result.GrossIncome,
	}
	return &response, nil
}
