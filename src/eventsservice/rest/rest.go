package rest

import (
	"goEventsSite/src/lib/msgqueue"
	"goEventsSite/src/lib/persistence"
	"net/http"

	"github.com/gorilla/mux"
)

// ServeApi uses gorilla/mux to serve the events API
func ServeApi(endpoint, tlsendpoint, tlscert, tlskey string, dbHandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) (chan error, chan error) {
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

	go func() {
		httpErrChan <- http.ListenAndServe(endpoint, r)
	}()
	// serve that shit up...securely
	go func() {
		httptlsErrChan <- http.ListenAndServeTLS(tlsendpoint, tlscert, tlskey, r)
	}()

	return httpErrChan, httptlsErrChan
}