package rest

import (
	"goEventsSite/src/lib/msgqueue"
	"goEventsSite/src/lib/persistence"
	"net/http"

	"github.com/gorilla/mux"
)

// ServeHTTP uses Gorilla/Mux to serve up the booking api
func ServeHTTP(endpoint, tlsendpoint string, dbHandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) (chan error, chan error) {
	handler := newBookingHandler(dbHandler, eventEmitter)
	r := mux.NewRouter()

	// define the booking subrouter
	bookingrouter := r.Path("/events/{eventID}/bookings").Subrouter()

	// list all bookings
	bookingrouter.
		Methods("GET").
		Path("").
		HandlerFunc(handler.getAllBookingsHandler)

	// add a new booking
	bookingrouter.
		Methods("POST").
		Path("").
		HandlerFunc(handler.addNewBookingHandler)

	// run blocking functions as goroutines
	httpErrChan := make(chan error)
	httptlsErrChan := make(chan error)

	go func() {
		httpErrChan <- http.ListenAndServe(endpoint, r)
	}()
	// serve that shit up...securely
	go func() {
		httptlsErrChan <- http.ListenAndServeTLS(tlsendpoint, "../../../etc/keys/cert.pem", "../../../etc/keys/key.pem", r)
	}()

	return httpErrChan, httptlsErrChan
}
