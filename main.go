package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	serv := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	log.Fatal(serv.ListenAndServe())
}
