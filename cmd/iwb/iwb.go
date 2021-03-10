package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dannywolfmx/iwb/api"
)

func main() {
	router := api.CreateHTTPServer()
	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Running the server in: %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
