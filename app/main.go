package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := "0.0.0.0:3001"
	log.Printf("Starting server: %s", addr)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("hit")
		fmt.Fprintln(w, "hello there")
	})

	log.Fatal(http.ListenAndServe(addr, mux))

}
