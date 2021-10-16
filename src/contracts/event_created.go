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
