package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	fmt.Println("Server Started.")
	// Passing mux as `http.Handler`
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal(err)
	}
}

/*
http.Handler Interface:
	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}
*/