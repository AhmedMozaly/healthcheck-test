package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		rw.Write([]byte("I'm healthy!!"))
	})
	mux.HandleFunc("/healthz", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(500)
		rw.Write([]byte("I'm not healthy :("))
	})
	mux.HandleFunc("/healthzz", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(503)
		rw.Write([]byte("I'm really not healthy :(((("))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "6294"
	}
	log.Printf("starting server on port %s", port)
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: mux,
	}

	log.Print(server.ListenAndServe())
}
