// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0--rc1
// source: movies.proto

package protogen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MovieInterface_SearchMovie_FullMethodName     = "/grpc.MovieInterface/SearchMovie"
	MovieInterface_MovieNowPlaying_FullMethodName = "/grpc.MovieInterface/MovieNowPlaying"
)

// MovieInterfaceClient is the client API for MovieInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieInterfaceClient interface {
	SearchMovie(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	MovieNowPlaying(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (MovieInterface_MovieNowPlayingClient, error)
}

type movieInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieInterfaceClient(cc grpc.ClientConnInterface) MovieInterfaceClient {
	return &movieInterfaceClient{cc}
}

func (c *movieInterfaceClient) SearchMovie(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, MovieInterface_SearchMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieInterfaceClient) MovieNowPlaying(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (MovieInterface_MovieNowPlayingClient, error) {
	stream, err := c.cc.NewStream(ctx, &MovieInterface_ServiceDesc.Streams[0], MovieInterface_MovieNowPlaying_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &movieInterfaceMovieNowPlayingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MovieInterface_MovieNowPlayingClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type movieInterfaceMovieNowPlayingClient struct {
	grpc.ClientStream
}

func (x *movieInterfaceMovieNowPlayingClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MovieInterfaceServer is the server API for MovieInterface service.
// All implementations must embed UnimplementedMovieInterfaceServer
// for forward compatibility
type MovieInterfaceServer interface {
	SearchMovie(context.Context, *Request) (*Response, error)
	MovieNowPlaying(*PageRequest, MovieInterface_MovieNowPlayingServer) error
	mustEmbedUnimplementedMovieInterfaceServer()
}

// UnimplementedMovieInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedMovieInterfaceServer struct {
}

func (UnimplementedMovieInterfaceServer) SearchMovie(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMovie not implemented")
}
func (UnimplementedMovieInterfaceServer) MovieNowPlaying(*PageRequest, MovieInterface_MovieNowPlayingServer) error {
	return status.Errorf(codes.Unimplemented, "method MovieNowPlaying not implemented")
}
func (UnimplementedMovieInterfaceServer) mustEmbedUnimplementedMovieInterfaceServer() {}

// UnsafeMovieInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieInterfaceServer will
// result in compilation errors.
type UnsafeMovieInterfaceServer interface {
	mustEmbedUnimplementedMovieInterfaceServer()
}

func RegisterMovieInterfaceServer(s grpc.ServiceRegistrar, srv MovieInterfaceServer) {
	s.RegisterService(&MovieInterface_ServiceDesc, srv)
}

func _MovieInterface_SearchMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieInterfaceServer).SearchMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieInterface_SearchMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieInterfaceServer).SearchMovie(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieInterface_MovieNowPlaying_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MovieInterfaceServer).MovieNowPlaying(m, &movieInterfaceMovieNowPlayingServer{stream})
}

type MovieInterface_MovieNowPlayingServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type movieInterfaceMovieNowPlayingServer struct {
	grpc.ServerStream
}

func (x *movieInterfaceMovieNowPlayingServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

// MovieInterface_ServiceDesc is the grpc.ServiceDesc for MovieInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.MovieInterface",
	HandlerType: (*MovieInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchMovie",
			Handler:    _MovieInterface_SearchMovie_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MovieNowPlaying",
			Handler:       _MovieInterface_MovieNowPlaying_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "movies.proto",
}
