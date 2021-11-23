package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(&InMemoryPlayerStore{})
	handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
