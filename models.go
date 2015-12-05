package main

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDatabase struct {
	s *mgo.Session
}

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
