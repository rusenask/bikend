package main

import (
	"os"
)

// Initial structure of configuration
type Configuration struct {
	mongoAddress string
	databaseName string
	ESRIEndpoint string
}

// AppConfig stores application configuration
var AppConfig Configuration

func initSettings() {

	// getting MongoDB connection details
	mongoAddress := os.Getenv("MongoURI")
	if mongoAddress == "" {
		mongoAddress = "localhost:27017"
	}
	AppConfig.mongoAddress = mongoAddress

	// getting default database
	defaultDB := os.Getenv("DatabaseName")
	if defaultDB == "" {
		defaultDB = "BikendDB"
	}
	AppConfig.databaseName = defaultDB

	esri := os.Getenv("ESRI")
	AppConfig.ESRIEndpoint = esri

}
