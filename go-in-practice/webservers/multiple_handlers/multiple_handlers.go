package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye/", goodbye)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Robel Yemane"
	}
	fmt.Fprint(res, "Hello, my name is", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Robel Yemane"
	}
	fmt.Fprint(res, "Goodbye ", name)
	fmt.Printf("Path: %s \n", req.URL.Path)
	fmt.Printf("Parts: %s", parts)
	fmt.Printf("Partsis:: %s", parts[0])

}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprintln(res, "The homepage")
	fmt.Printf("Req: %s %s\n", req.Host, req.URL.Path)
}
