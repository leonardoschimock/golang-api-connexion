package main

import (
	"extensao-api/router"
	"log"
	"net/http"
)

func main() {
	r := router.New()
	const addr = ":8080"
	log.Printf("Starting server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
