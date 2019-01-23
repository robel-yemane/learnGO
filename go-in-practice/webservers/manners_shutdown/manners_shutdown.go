package main

import (
	"fmt"
	"github.com/braintree/manners"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	handler := newHandler() // gets instance of a handler

	//Sets up monitoring of operating system signals
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)

	//Starts the web server
	manners.ListenAndServe(":8080", handler)
}

func newHandler() *handler {
	return &handler{}
}

// Handler responding to requets

type handler struct{}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Robel Yemane"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

//Waits for shutdown signal and reacts
func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}
