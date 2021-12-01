package rest

import (
	"encoding/json"
	"fmt"
	"goEventsSite/src/contracts"
	"goEventsSite/src/lib/msgqueue"
	"goEventsSite/src/lib/persistence"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// define the booking handler structure
type bookingHandler struct {
	dbhandler    persistence.DatabaseHandler
	eventEmitter msgqueue.EventEmitter
}

func newBookingHandler(dbhandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) *bookingHandler {
	return &bookingHandler{
		dbhandler:    dbhandler,
		eventEmitter: eventEmitter,
	}
}

// addNewBookingHandler adds a new booking for a user to the db
func (bh *bookingHandler) addNewBookingHandler(w http.ResponseWriter, r *http.Request) {
	routeVars := mux.Vars(r)
	// get the fucking eventid from the mux route variable eventid
	eventID, ok := routeVars["eventID"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{error: missing eventid}`)
		return
	}

	// find the god damn event that matches the fucking eventid
	event, err := bh.dbhandler.FindEvent([]byte(eventID))
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, `{error: event %s could not be found: %s}`, eventID, err)
	}

	// get the booking request json data from the body and decode that shit
	var bookingRequest persistence.Booking
	err = json.NewDecoder(r.Body).Decode(&bookingRequest)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: could not decode booking json data %s}`, err)
	}
	if seats := bookingRequest.Seats; seats <= 0 { // cant have less than one fucking booking now can you???
		w.WriteHeader(400)
		fmt.Fprintf(w, `{error: seat number must be greater than %s}`, err)
		return
	}

	booking := persistence.Booking{
		Date:    time.Now().Unix(),
		EventID: []byte(eventID),
		Seats:   bookingRequest.Seats,
	}

	// add the booking to the db instance
	err = bh.dbhandler.AddBooking([]byte(eventID), booking)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: could not add booking data to database %s}`, err)
	}

	// create the AMQP message and send to the broker
	msg := contracts.EventBookedEvent{
		EventID: event.ID.Hex(),
		UserID:  "someUserID",
	}
	bh.eventEmitter.Emit(&msg)

	// write the response by replying with the booking
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	json.NewEncoder(w).Encode(&booking)
}
