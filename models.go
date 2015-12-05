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
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Rating      float64       `json:"rating"`
	Description string        `json:"description"`
	Author      bson.ObjectId `json:"author"`
}

type User struct {
	Id            bson.ObjectId  `json:"id" bson:"_id"` // user id
	HostingPlaces []HostingPlace `json:"hostingPlaces"` // hosting places that this user has registered
	BikeLocation  HostingPlace   `json:"bikeLocation"`  // where my bike is now (empty if you haven't put your bike)
	Reviews       []Review       `json:"reviews"`
	UserID        string         `json:"userID"`
	ProfilePicUrl string         `json:"profilePic"`
	FirstName     string         `json:"firstName"`
	LastName      string         `json:"lastName"`
	Gender        string         `json:"gender"`
}

type Booking struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	Date         time.Time     `json:"time"`         // when this booking happened
	User         bson.ObjectId `json:"user"`         // who did the booking
	Host         bson.ObjectId `json:"host"`         // who's owner
	HostingPlace bson.ObjectId `json:"hostingPlace"` // where is this booking taking place
}

type HostingPlace struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Host     bson.ObjectId `json:"host"`     // who is hosting this place
	Space    int           `json:"space"`    // how many bikes can you put here
	Active   bool          `json:"active"`   // is it active or not
	Long     float64       `json:"long"`     // longitude
	Lat      float64       `json:"lat"`      // latitude
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

	err := c.Find(bson.M{"userID": userID}).One(&result)

	return result, err
}

func (db *MongoDatabase) addUser(user User) error {
	c := db.s.DB(AppConfig.databaseName).C(user_collection)
	id := bson.NewObjectId()

	log.WithFields(log.Fields{
		"bsonID": id,
	}).Info("ID for document")

	user.Id = id

	err := c.Insert(user)

	return err
}
