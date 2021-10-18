package kafka

// messageEnvelope defines the payload structure for a kafka message
type messageEnvelope struct {
	EventName string      `json:"eventName"`
	Payload   interface{} `json:"payload"`
}
