package main

import (
	"log"
	"net/http"
	"time"

	"danirod.es/pkg/iwb/api"
)

func main() {
	router := api.CreateHTTPServer()
	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
