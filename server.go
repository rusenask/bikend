package main

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

func getBoneRouter(h HTTPClientHandler) *bone.Mux {
	mux := bone.New()
	// add new users
	mux.Post("/api/users", http.HandlerFunc(h.addUserHandler))
	mux.Get("/api/users", http.HandlerFunc(h.getAllUsersHandler))
	// returns user and his/her bike store locations and also where he booked
	mux.Get("/api/users/:user", http.HandlerFunc(h.getUserHandler))


	mux.Handle("/*", http.FileServer(http.Dir("static/dist")))

	return mux
}