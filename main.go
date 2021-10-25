package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		rw.Write([]byte("hello world!"))
	})
	mux.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		rw.Write([]byte("I'm healthy!!"))
	})
	mux.HandleFunc("/healthz", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(500)
		rw.Write([]byte("I'm not healthy :("))
	})

	log.Print("starting server on port 9999")
	server := &http.Server{
		Addr:    "0.0.0.0:9999",
		Handler: mux,
	}

	log.Print(server.ListenAndServe())
}
