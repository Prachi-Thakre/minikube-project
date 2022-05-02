package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>PONG</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ping", pong)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}