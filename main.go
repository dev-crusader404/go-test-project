package main

import (
	"flag"
	"fmt"
	"net/http"

	grpsv "github.com/dev-crusader404/go-test-project/grpc/server"
	sv "github.com/dev-crusader404/go-test-project/restapi"
	"github.com/dev-crusader404/go-test-project/startup"
)

var (
	logger       = sv.Logger
	makeHTTPFunc = sv.MakeHTTPFunc
	grpcAddr     = flag.String("grpc", ":5001", "listen address of the grpc transport")
)

func main() {

	startup.Load()
	flag.Parse()
	fmt.Printf("\nGRPC server running on port %s", *grpcAddr)
	grpsv.RunGRPCServer(*grpcAddr)

	s := sv.NewDB()
	http.HandleFunc("/", logger(makeHTTPFunc(s, sv.Handler)))
	http.ListenAndServe(":8080", nil)
}
