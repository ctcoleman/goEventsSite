package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"goEventsSite/src/contracts"
	"goEventsSite/src/lib/msgqueue"
	"goEventsSite/src/lib/persistence"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type eventHandler struct {
	dbhandler    persistence.DatabaseHandler
	eventEmitter msgqueue.EventEmitter
}

// newEventHandler creates a new session to handlers
func newEventHandler(databaseHandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) *eventHandler {
	return &eventHandler{
		dbhandler:    databaseHandler,
		eventEmitter: eventEmitter,
	}
}

// findEventHandler gets an event from a db instance by a given search criteria
func (eh *eventHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {
	var event persistence.Event
	var err error

	vars := mux.Vars(r)
	// validate search criteria
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search criteria found. Either search by id or name}`)
		return
	}

	// validate search key
	searchkey, ok := vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search keys found. Either search by id or name.}`)
		return
	}

	// find by name or id
	switch strings.ToLower(criteria) {
	case "name":
		event, err = eh.dbhandler.FindEventByName(searchkey)
	case "id":
		id, err := hex.DecodeString(searchkey)
		if err == nil {
			event, _ = eh.dbhandler.FindEvent(id)
		}
	}
	if err != nil {
		fmt.Fprintf(w, `{error %s}`, err)
		return
	}

	// encode to json
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&event)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Unable to encode json %s}`, err)
	}
}

// getAllEventsHandler gets all events stored in db instance
func (eh *eventHandler) getAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := eh.dbhandler.FindAllAvailableEvents()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Could not find all events in db %s}`, err)
		return
	}

	// encode gathered events to json
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Unable to encode event json data %s}`, err)
	}
}

func (eh *eventHandler) getAllLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := eh.dbhandler.FindAllLocations()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Could not find all locations in db: %s}`, err)
	}

	// encode gathered locations to json
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&locations)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Unable to encode location json data %s}`, err)
	}
}

func (eh *eventHandler) getEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventID, ok := vars["eventID"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{error: Invalid eventID}`)
		return
	}

	id, err := hex.DecodeString(eventID)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Could not decode eventID: %s}`, err)
	}

	event, err := eh.dbhandler.FindEvent(id)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Could not find event with eventID %s: %s}`, eventID, err)
	}

	// encode gathered event data to json
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&event)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Unable to encode event json data: %s`, err)
	}

}

// newEventHandler creates a new event based on json body in the request
func (eh *eventHandler) newEventHandler(w http.ResponseWriter, r *http.Request) {
	var event persistence.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Could not decode event data %s}`, err)
		return
	}

	// add event to db
	id, err := eh.dbhandler.AddEvent(event)
	if nil != err {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Unable add new event %d %s}`, id, err)
		return
	}

	// send out EventCreatedEvent AMQP message
	msg := contracts.EventCreatedEvent{
		ID:         hex.EncodeToString(id),
		Name:       event.Name,
		LocationID: string(event.Location.ID),
		Start:      time.Unix(event.StartDate, 0),
		End:        time.Unix(event.EndDate, 0),
	}
	eh.eventEmitter.Emit(&msg)
}

// newLocation creates a new location based on body json data
func (eh *eventHandler) newLocation(w http.ResponseWriter, r *http.Request) {
	var location persistence.Location
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Could not decode location data %s}`, err)
		return
	}

	// add the location to the database
	location, err = eh.dbhandler.AddLocation(location)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Could not add location to db: %s}`, err)
	}

	// send out LocationCreatedEvent AMQP message
	msg := contracts.LocationCreatedEvent{
		ID:     location.ID.Hex(),
		Name:   location.Name,
		City:   location.City,
		State:  location.State,
		Venues: location.Venues,
	}
	eh.eventEmitter.Emit(&msg)
}
