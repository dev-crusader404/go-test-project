package main

import (
	"net/http"

	sv "github.com/dev-crusader404/go-test-project/restapi"
	"github.com/dev-crusader404/go-test-project/startup"
)

var (
	logger       = sv.Logger
	makeHTTPFunc = sv.MakeHTTPFunc
)

func main() {
	// val.RunCreditCardValidator()
	startup.Load()

	s := sv.NewDB()
	http.HandleFunc("/", logger(makeHTTPFunc(s, sv.Handler)))
	http.ListenAndServe(":8080", nil)
}
