package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

var users = []User{}

func (s *api) usersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		s.getUsersHandler(w, r)
	case http.MethodPost:
		s.createUsersHandler(w, r)
	default:
	}

}

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
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	if err := insertUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func insertUser(pUser User) error {

	if pUser.Email == "" {
		return errors.New("email is required")
	}

	if pUser.FirstName == "" {
		return errors.New("first name is required")
	}

	for _, user := range users {

		if user.Email == pUser.Email {
			return errors.New("user already exists")
		}
	}

	users = append(users, pUser)
	return nil
}
