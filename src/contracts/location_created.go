package contracts

import "goEventsSite/src/lib/persistence"

// details of AMQP message sent whenever another fucking location is created
type LocationCreatedEvent struct {
	ID     string              `json:"id"`
	Name   string              `json:"name"`
	City   string              `json:"city"`
	State  string              `json:"state"`
	Venues []persistence.Venue `json:"venue"`
}

// generate a fucking AMQP event name for when a location is created
func (e *LocationCreatedEvent) EventName() string {
	return "locationCreated"
}

// generate a god damn Partition Key because kafka needs more god damn details for each Event
func (e *LocationCreatedEvent) PartitionKey() string {
	return e.ID // you already have a fucking unique ID you can use dumb ass
}
