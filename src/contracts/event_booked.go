package contracts

// EventBookedEvent describes the AMQP message details sent each time an event is booked
type EventBookedEvent struct {
	EventID string `json:"event_id"`
	UserID  string `json:"user_id"`
}

// generate a fucking AMQP event name for each time an event is booked
func (e *EventBookedEvent) EventName() string {
	return "eventBooked"
}
