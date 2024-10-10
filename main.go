package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/dev-crusader404/go-test-project/client"
	grpsv "github.com/dev-crusader404/go-test-project/grpc/server"
	mv "github.com/dev-crusader404/go-test-project/internal"
	sv "github.com/dev-crusader404/go-test-project/restapi"
	"github.com/dev-crusader404/go-test-project/startup"
)

var (
	logger          = sv.Logger
	fetcherHTTPFunc = sv.Fetcher
	grpcAddr        = flag.String("grpc", ":5001", "listen address of the grpc transport")
)

func main() {

	startup.Load()
	flag.Parse()
	fmt.Printf("\nGRPC server running on port %s\n", *grpcAddr)
	go grpsv.RunGRPCServer(*grpcAddr)
	m := &mv.Movie{Client: client.GetClient()}
	http.HandleFunc("/", logger(fetcherHTTPFunc(m, sv.GetMovieHandler)))
	fmt.Println("Http server running on port :8080")
	http.ListenAndServe(":8080", nil)
}
