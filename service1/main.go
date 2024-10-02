package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("*************************************")
	fmt.Println("Server is running on :8081...")
	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}