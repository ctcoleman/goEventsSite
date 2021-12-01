package rest

import (
	"fmt"
	"goEventsSite/src/lib/msgqueue"
	"goEventsSite/src/lib/persistence"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// ServeHTTP uses Gorilla/Mux to serve up the booking api
func ServeAPI(endpoint, tlsendpoint, tlscert, tlskey string, dbHandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) (chan error, chan error) {
	handler := newBookingHandler(dbHandler, eventEmitter)
	r := mux.NewRouter()

	// add a new booking
	r.
		Methods("POST").
		Path("/events/{eventID}/bookings").
		HandlerFunc(handler.addNewBookingHandler)

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
