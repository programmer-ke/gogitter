package main

import (
	"log"

	"github.com/programmer-ke/gogitter/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
