package main

import (
	"log"
	"net/http"
)

func main() {
	api := &api{addr: ":8085"}

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("/users", api.usersHandler)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
