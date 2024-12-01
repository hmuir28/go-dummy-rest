package main

import (
	"encoding/json"
	"net/http"
)

type api struct {
	addr string
}

var users = []User{}

func (s *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (s *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	var payload User

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := User{
		firstName: payload.firstName,
		lastName:  payload.lastName,
	}

	users = append(users, user)

	w.WriteHeader(http.StatusCreated)
}
