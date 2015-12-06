package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDatabase struct {
	s *mgo.Session
}

// Constants representing collection names
const user_collection string = "users"
const review_collection string = "reviews"
const booking_collection string = "bookings"
const places_collection string = "places"

type Review struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Rating      float64       `json:"rating"`
	Description string        `json:"description"`
	Author      bson.ObjectId `json:"author,omitempty"`
}

type User struct {
	Id            bson.ObjectId  `json:"id" bson:"_id,omitempty"`                                // user id
	HostingPlaces []HostingPlace `json:"hostingPlaces,omitempty" bson:"hostingPlaces,omitempty"` // hosting places that this user has registered
	BikeLocation  HostingPlace   `json:"bikeLocation,omitempty" bson:"bikeLocation,omitempty"`   // where my bike is now (empty if you haven't put your bike)
	Reviews       []Review       `json:"reviews,omitempty" bson:"reviews,omitempty"`
	UserID        string         `json:"userID"` // user email
	ProfilePicUrl string         `json:"profilePic"`
	FirstName     string         `json:"firstName"`
	LastName      string         `json:"lastName"`
	Gender        string         `json:"gender,omitempty"`
}

type Booking struct {
	Id   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Date time.Time     `json:"time"`           // when this booking happened
	User string        `json:"user,omitempty"` // who did the booking
	Host string        `json:"host,omitempty"` // who's owner
	Long float64       `json:"long"`           // longitude
	Lat  float64       `json:"lat"`            // latitude
}

type HostingPlace struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Host     string        `json:"host"`     // who is hosting this place (email address)
	Space    int           `json:"space"`    // how many bikes can you put here
	Active   bool          `json:"active"`   // is it active or not
	Long     float64       `json:"long"`     // longitude
	Lat      float64       `json:"lat"`      // latitude
	Address  string        `json:"address"`  // address
	Bookings []Booking     `json:"bookings"` // current bookings
}

// filterCategories function searches in categories collection based on supplied keywords
func (db *MongoDatabase) getUsers() (results []User, err error) {
	c := db.s.DB(AppConfig.databaseName).C(user_collection)

	err = c.Find(nil).All(&results)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Warn("Category search failed.")
		return nil, err
	} else {
		return results, nil
	}

}

func (db *MongoDatabase) getUser(userID string) (User, error) {
	c := db.s.DB(AppConfig.databaseName).C(user_collection)

	var result User

	err := c.Find(bson.M{"userid": userID}).One(&result)

	// getting places
	pc := db.s.DB(AppConfig.databaseName).C(places_collection)
	var places []HostingPlace
	err = pc.Find(bson.M{"host": userID}).All(&places)

	if len(places) > 0 {
		place := places[0]
		// bookings
		bc := db.s.DB(AppConfig.databaseName).C(booking_collection)
		var bookings []Booking

		err = bc.Find(bson.M{
			"host": userID,
			"lat":  place.Lat,
			"long": place.Long,
		}).All(&bookings)

		result.HostingPlaces = places

		place.Bookings = bookings
		result.HostingPlaces[0] = place
		return result, err
	} else {
		return result, err

	}

}

func (db *MongoDatabase) addUser(user User) error {
	c := db.s.DB(AppConfig.databaseName).C(user_collection)
	id := bson.NewObjectId()

	log.WithFields(log.Fields{
		"bsonID": id,
	}).Info("ID for user document")

	user.Id = id

	err := c.Insert(user)

	return err
}

func (db *MongoDatabase) addHostingPlace(place HostingPlace) error {

	// adding place
	c := db.s.DB(AppConfig.databaseName).C(places_collection)
	id := bson.NewObjectId()
	log.WithFields(log.Fields{
		"bsonID": id,
	}).Info("ID for hosting place document")
	place.Id = id

	err := c.Insert(place)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Got error while adding new place.")
	}

	return err
}

func (db *MongoDatabase) addBooking(booking Booking) error {
	// adding place
	c := db.s.DB(AppConfig.databaseName).C(booking_collection)
	id := bson.NewObjectId()
	log.WithFields(log.Fields{
		"bsonID": id,
	}).Info("ID for booking document")
	booking.Id = id
	booking.Date = time.Now()

	err := c.Insert(booking)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Got error while adding new booking.")
	}

	return err
}
