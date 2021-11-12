package rest

import (
	"fmt"
	"goEventsSite/src/lib/msgqueue"
	"goEventsSite/src/lib/persistence"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// ServeApi uses gorilla/mux to serve the events API
func ServeAPI(endpoint, tlsendpoint, tlscert, tlskey string, dbHandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) (chan error, chan error) {
	handler := newEventHandler(dbHandler, eventEmitter)
	r := mux.NewRouter()

	// define the events subrouter
	eventsrouter := r.PathPrefix("/events").Subrouter()

	// list all events
	eventsrouter.
		Methods("GET").
		Path("").
		HandlerFunc(handler.getAllEventsHandler)

	// create a new event
	eventsrouter.
		Methods("POST").
		Path("").
		HandlerFunc(handler.newEventHandler)

	// go to an event
	eventsrouter.
		Methods("GET").
		Path("/{eventID}").HandlerFunc(handler.getEventHandler)

	// searching for events
	eventsrouter.
		Methods("GET").
		Path("/{SearchCriteria}/{search}").
		HandlerFunc(handler.findEventHandler)

	// define the locations subrouter
	locationrouter := r.PathPrefix("/locations").Subrouter()

	// add location
	locationrouter.
		Methods("POST").
		Path("/locations").
		HandlerFunc(handler.newLocation)

	// get all locations
	locationrouter.
		Methods("GET").
		Path("/locations").
		HandlerFunc(handler.getAllLocations)

	// run blocking functions as goroutines
	httpErrChan := make(chan error)
	httptlsErrChan := make(chan error)

	// enable cors to allow http requests from frontend
	server := handlers.CORS()(r)

	// serve that shit up
	go func() {
		fmt.Println("restful http server listening at " + endpoint)
		httpErrChan <- http.ListenAndServe(endpoint, server)
	}()
	// serve that shit up...securely
	go func() {
		fmt.Println("secure restful https server listening at " + tlsendpoint)
		httptlsErrChan <- http.ListenAndServeTLS(tlsendpoint, tlscert, tlskey, server)
	}()

	return httpErrChan, httptlsErrChan
}
