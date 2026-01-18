package main

import (
	"log"
	"net/http"

	"ponglockdown/server/api"
)

func main() {
	mux := http.NewServeMux()
	api.Register(mux)
	addr := ":8080"
	log.Println("AAA Lockdown server listening on", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
