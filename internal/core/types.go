package core

// EventType represents the type of a given event.
type EventType string

// Event represents an event with a type and payload.
type Event struct {
	Type    EventType
	Payload interface{}
}

// Listener represents a listener for a given event type.
type Listener struct {
	ID      string
	Module  string
	Handler func(e Event)
}
