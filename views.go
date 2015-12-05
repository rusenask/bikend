package main

import (
	"net/http"
)

func (h *HTTPClientHandler) addUserHandler (w http.ResponseWriter, r *http.Request) {
	// adding new user to database
}

func (h *HTTPClientHandler) getAllUsersHandler (w http.ResponseWriter, r *http.Request) {
	// displaying all users
}

func (h *HTTPClientHandler) getUserHandler (w http.ResponseWriter, r *http.Request) {
	// display current users locations for HOSTING and where he will be parking or is parking
}

