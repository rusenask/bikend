package main

import (
	"gopkg.in/mgo.v2"
)

type MongoDatabase struct {
	s *mgo.Session
}