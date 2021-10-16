package contracts

import "time"

// EventCreatedEvent describes the AMQP message details each time a new booking event is create
type EventCreatedEvent struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	LocationID string    `json:"location_id"`
	Start      time.Time `json:"start_time"`
	End        time.Time `json:"end_time"`
}

// generate a fucking AMQP event name for when a new booking event is created
func (e *EventCreatedEvent) EventName() string {
	return "eventCreated"
}

// generate a god damn Partition Key because kafka needs more god damn details for each Event
func (e *EventCreatedEvent) PartitionKey() string {
	return e.ID // you already have a fucking unique ID you can use dumb ass
}
