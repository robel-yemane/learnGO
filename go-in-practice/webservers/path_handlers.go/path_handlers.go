package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

func main() {
	pr := newPathResolver() // gets an instance of a path-based router
	//maps functions to paths
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye)
	http.ListenAndServe(":8080", pr)
}

// creates new initialised pathresolver
func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

//Adds paths to internal lookup
func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//Constructs our method + path to check
	check := req.Method + " " + req.URL.Path
	// Iterates over registered paths
	for pattern, handlerFunc := range p.handlers {
		//checks whether current path matches a registered one
		//strict checking takes place here
		if ok, err := path.Match(pattern, check); ok && err == nil {
			//Executes the handler function for a matched path
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}
	http.NotFound(res, req)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "robel yemane"
	}
	fmt.Fprint(res, "hello, my name is ", name)
}
func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "robel yemane"
	}
	fmt.Fprint(res, "Goodbye ", name)
}
