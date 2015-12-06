package main

import (
	log "github.com/Sirupsen/logrus"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UsersResource struct {
	Data []User `json:"data"`
}

type UserResource struct {
	Data User `json:"data"`
}

// Structure representing error
type errorResponse struct {
	Msg string `json:"msg"`
}

var ServerName = "http://localhost:8080"

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
	w.Header().Add("Access-Control-Allow-Origin", ServerName)
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

	userid, _ := r.URL.Query()["q"]
	// looking for specific user
	if len(userid) > 0 {
		log.WithFields(log.Fields{
			"userid": userid[0],
		}).Info("Looking for user..")

		user, err := h.db.getUser(userid[0])

		if err == nil {
			// Marshal provided interface into JSON structure
			response := UserResource{Data: user}
			uj, _ := json.Marshal(response)

			// Write content-type, statuscode, payload
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			fmt.Fprintf(w, "%s", uj)
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
			return

		}
	}

	log.Warn(len(userid))
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
	response := UsersResource{Data: results}
	uj, _ := json.Marshal(response)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// addPlaceHandler add new hosting place, provide json
func (h *HTTPClientHandler) addPlaceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", ServerName)
	// adding new hosting place to database
	var hostingPlaceRequest HostingPlace
	log.Info("adding place........")
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

	err = json.Unmarshal(body, &hostingPlaceRequest)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Failed to unmarshal json!")
	}

	log.WithFields(log.Fields{
		"body":   string(body),
		"host":   hostingPlaceRequest.Host,
		"active": hostingPlaceRequest.Active,
		"lat":    hostingPlaceRequest.Lat,
		"long":   hostingPlaceRequest.Long,
	}).Info("Got place info")

	err = h.db.addHostingPlace(hostingPlaceRequest)

	if err == nil {
		// adding it to esri
		h.addEsriNode(hostingPlaceRequest)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(201) // place inserted
		return
	} else {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Failed to insert hosting place..")

		content, code := responseDetailsFromMongoError(err)

		// Marshal provided interface into JSON structure
		uj, _ := json.Marshal(content)

		// Write content-type, statuscode, payload
		writeJsonResponse(w, &uj, code)

	}

}

func (h *HTTPClientHandler) getPlaceHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPClientHandler) addBookingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", ServerName)
	// adding new hosting place to database
	var bookingRequest Booking
	log.Info("adding place........")
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

	err = json.Unmarshal(body, &bookingRequest)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Failed to unmarshal json!")
	}

	log.WithFields(log.Fields{
		"body": string(body),
		"host": bookingRequest.Host,
		"user": bookingRequest.User,
		"lat":  bookingRequest.Lat,
		"long": bookingRequest.Long,
	}).Info("Got place info")

	err = h.db.addBooking(bookingRequest)

	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(201) // booking inserted
		return
	} else {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Failed to insert booking, fork it..")

		content, code := responseDetailsFromMongoError(err)

		// Marshal provided interface into JSON structure
		uj, _ := json.Marshal(content)

		// Write content-type, statuscode, payload
		writeJsonResponse(w, &uj, code)

	}
}
