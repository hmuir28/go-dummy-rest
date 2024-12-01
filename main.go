package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (s *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Retrieved user"))
}

func (s *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created user"))
}

func main() {
	api := &api{addr: ":8085"}

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	err := http.ListenAndServe(server.Addr, server.Handler)

	if err != nil {
		log.Fatal(err)
	}
}
