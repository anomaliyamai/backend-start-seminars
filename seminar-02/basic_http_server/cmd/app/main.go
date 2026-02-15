package main

import (
	"flag"
	"http_server/http"
	"http_server/storage"
	"log"
)

func main() {
	addr := flag.String("addr", ":8080", "address for http server")

	s := storage.NewRaiStorage()

	log.Printf("Starting server on %s", *addr)
	if err := http.CreateAndRunServer(s, *addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
