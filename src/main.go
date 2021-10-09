package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello world</h1>")
	fmt.Println("Endpoint hit: homepage")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>health check</h1>")
	fmt.Println("Endpoint hit: checkpage")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/health_check", check)
	fmt.Println("Server starting ....")
	http.ListenAndServe(":8181", nil)
}
