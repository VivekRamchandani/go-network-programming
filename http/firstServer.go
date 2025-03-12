package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL: %s\n", r.URL)
	fmt.Printf("URL Query: %s", r.URL.Query())
	fmt.Fprintf(w, "Hello, %s", r.URL)
}

func main() {
	// HandleFunc registers the handler function for the given pattern in [DefaultServeMux].
	http.HandleFunc("/", handler)
	fmt.Println("Sever Started")
	// Passing `nil` as handler(http.Handler) will use [DefaultServerMux]
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
