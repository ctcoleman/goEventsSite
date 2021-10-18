package listener

import (
	"fmt"
	"goEventsSite/src/contracts"
	"goEventsSite/src/lib/msgqueue"
	"goEventsSite/src/lib/persistence"
	"log"

	"gopkg.in/mgo.v2/bson"
)

type EventProcessor struct {
	EventListener msgqueue.EventListener
	Database      persistence.DatabaseHandler
}

// ProcessEvents method listens for eventCreated event and processes event through handleEvent function
func (p *EventProcessor) ProcessEvents() error {
	log.Println("Booking service listening to events...")

	received, errors, err := p.EventListener.Listen("eventCreated")
	if err != nil {
		panic(err)
	}

	for {
		select {
		case evt := <-received:
			fmt.Printf("event received %T: %s\n", evt, evt)
			p.handleEvent(evt)
		case err = <-errors:
			log.Printf("-- errror while processing event -- %s\n", err)
		}
	}
}

// handleEvent method calls handler based on event
func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	// if event is created then fucking update the database
	case *contracts.EventCreatedEvent:
		p.Database.AddEvent(persistence.Event{ID: bson.ObjectId(e.ID)})
		if !bson.IsObjectIdHex(e.ID) {
			log.Printf("event %v did not contain valid object ID", e)
		}
		log.Printf("event %s created: %s", e.ID, e)
	// if location is created then fucking update the database
	case *contracts.LocationCreatedEvent:
		p.Database.AddLocation(persistence.Location{ID: bson.ObjectId(e.ID)})
		log.Printf("location %s created: %v", e.ID, e)
	// idk wtf you are trying to do but it's not a fucking event
	default:
		log.Printf("unknown event: %t", e)
	}
}
