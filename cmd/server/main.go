package main

import (
	"github.com/dmorenoh/proglog/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatalf("error starting server: %v", srv.ListenAndServe())
}
