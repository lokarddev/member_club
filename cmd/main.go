package main

import (
	"fmt"
	"log"
	"member_club/internal"
	"net/http"
)

const (
	port = 8000
)

func main() {
	log.Printf("Started server listening on port %v", port)
	internal.RegisterRoutes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
