package main

import (
	"flag"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/go-zoo/bone"
	"github.com/meatballhat/negroni-logrus"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
)

// Client structure to be injected into functions to perform HTTP calls
type Client struct {
	HTTPClient *http.Client
}

// HTTPClientHandler used for passing http client connection and template
// information back to handlers, mostly for testing purposes
type HTTPClientHandler struct {
	http Client
	r    *render.Render
	db   MongoDatabase
}

func main() {
	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)
	log.SetFormatter(&log.TextFormatter{})

	port := flag.String("port", ":8080", "application port")
	flag.Parse()

	// geting db settings
	initSettings()

	// getting Mongo connection
	// mongoAddress can be a list of master/slaves.
	// session, err := mgo.Dial("server1.example.com,server2.example.com")
	session, err := mgo.Dial(AppConfig.mongoAddress)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior. Monotonic consistency will start reading from
	// a slave if possible, so that the load is better distributed, and once the first write happens the
	// connection is switched to the master. This offers consistent reads and writes,
	// but may not show the most up-to-date data on reads which precede the first write.
	session.SetMode(mgo.Monotonic, true)

	// ensuring indexes for name and category keywords
	c := session.DB(AppConfig.databaseName).C(user_collection)
	index := mgo.Index{
		Key:        []string{"userid"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		log.WithFields(log.Fields{
			"Error": err.Error(),
		}).Error("Failed to ensure full-text search indexes for users collection!")
	} else {
		log.Info("Indexes for users collection ensured!")

		// app starting
		log.WithFields(log.Fields{
			"mongoAddress": AppConfig.mongoAddress,
			"databaseName": AppConfig.databaseName,
		}).Info("app is starting")

		// getting base template and handler struct
		r := render.New()

		// getting database struct
		database := MongoDatabase{s: session}

		h := HTTPClientHandler{http: Client{&http.Client{}},
			r:  r,
			db: database,
		}

		mux := getBoneRouter(h)
		n := negroni.Classic()
		n.Use(negronilogrus.NewMiddleware())
		n.UseHandler(mux)
		n.Run(*port)
	}
}

func getBoneRouter(h HTTPClientHandler) *bone.Mux {
	mux := bone.New()

	// add new users
	mux.Post("/api/users", http.HandlerFunc(h.addUserHandler))
	mux.Get("/api/users", http.HandlerFunc(h.getAllUsersHandler))

	// add place
	mux.Post("/api/places", http.HandlerFunc(h.addPlaceHandler))
	// get places
	mux.Get("/api/places", http.HandlerFunc(h.getPlaceHandler))

	mux.Handle("/*", http.FileServer(http.Dir("static/bikend/www")))

	return mux
}
