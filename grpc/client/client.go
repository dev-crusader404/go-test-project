package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	svhandler "github.com/dev-crusader404/go-test-project/grpc/protogen"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	reverseProxyAddr = flag.String("proxyserver", ":8081", "listen address of the proxy service")
)

func main() {
	// Set up a connection to the grpc server.
	flag.Parse()
	grpcServerAddr := os.Getenv("GRPC_SERVER_URL")
	if grpcServerAddr == "" {
		grpcServerAddr =
			"localhost:5001" // Fallback to localhost
	}
	fmt.Println("GRPC server is running on " + grpcServerAddr)
	conn, err := grpc.Dial(grpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to grpc service: %v", err)
	}
	defer conn.Close()
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	if err = svhandler.RegisterMovieInterfaceHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the order server: %v", err)
	}
	// start listening to requests from the gateway server
	fmt.Println("API gateway server is running on " + *reverseProxyAddr)
	if err = http.ListenAndServe(*reverseProxyAddr, mux); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}
}
