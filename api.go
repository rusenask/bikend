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

// Structure representing error
type errorResponse struct {
	Msg string `json:"msg"`
}

func responseDetailsFromMongoError(error interface{}) (content errorResponse, code int) {
	content = errorResponse{Msg: fmt.Sprint(error)}
	code = 400
	if content.Msg == "not found" {
		code = 404
	}
	return content, code
}

// write json response to http response
func writeJsonResponse(w http.ResponseWriter, content *[]byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(*content)
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
		"lastName":      userRequest.LastName,
		"userID":        userRequest.UserID,
		"profilePicUrl": userRequest.ProfilePicUrl,
		"gender":        userRequest.Gender,
		"body":          string(body),
	}).Info("Got user info")

	// adding user
	err = h.db.addUser(userRequest)

	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(201) // user inserted
		return
	} else {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Failed to insert..")

		content, code := responseDetailsFromMongoError(err)

		// Marshal provided interface into JSON structure
		uj, _ := json.Marshal(content)

		// Write content-type, statuscode, payload
		writeJsonResponse(w, &uj, code)

	}

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

	log.WithFields(log.Fields{
		"count": len(results),
	}).Info("number of users")

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
