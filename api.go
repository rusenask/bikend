package main

import (
	log "github.com/Sirupsen/logrus"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserResource struct {
	Data []User `json:"data"`
}

// addUserHandler used to add new user
func (h *HTTPClientHandler) addUserHandler(w http.ResponseWriter, r *http.Request) {
	// adding new user to database
	var userRequest User

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		// failed to read response body
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Could not read response body!")
		http.Error(w, "Failed to read request body.", 400)
		return
	}

	err = json.Unmarshal(body, &userRequest)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // can't process this entity
		return
	}
	log.WithFields(log.Fields{
		"firstName":     userRequest.FirstName,
		"firstName":     userRequest.LastName,
		"userID":        userRequest.UserID,
		"profilePicUrl": userRequest.ProfilePicUrl,
		"gender":        userRequest.Gender,
		"body":          string(body),
	}).Info("New user inserted!")

}

// getAllUsersHandler used to get all users
func (h *HTTPClientHandler) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	// displaying all users
	results, err := h.db.getUsers()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Got error when tried to get all users")
	}
	// Marshal provided interface into JSON structure
	response := UserResource{Data: results}
	uj, _ := json.Marshal(response)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (h *HTTPClientHandler) getUserHandler(w http.ResponseWriter, r *http.Request) {
	// display current users locations for HOSTING and where he will be parking or is parking
}

func (h *HTTPClientHandler) updateUserHandler(w http.ResponseWriter, r *http.Request) {

}
