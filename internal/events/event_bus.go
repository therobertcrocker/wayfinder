package events

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/therobertcrocker/wayfinder/internal/common/logging"
)

var (
	Log = logging.Log
)

// EventBus represents the event bus.
type EventManager interface {
	Subscribe(eventType EventType, listener *Listener) error
	Unsubscribe(eventType EventType, listener *Listener) error
	Publish(e Event) error
}

type EventBus struct {
	subscribers map[EventType][]*Listener
}

// NewEventBus creates a new event bus.
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[EventType][]*Listener),
	}
}

// Subscribe subscribes a listener to a given event type.
func (b *EventBus) Subscribe(eventType EventType, listener *Listener) error {

	// Check if listener is valid
	err := isValidListener(listener)
	if err != nil {
		Log.WithFields(logrus.Fields{"module": listener.Module, "event": eventType}).Warnf("listener is not valid: %v", err)
		return err
	}

	// add listener to subscribers
	b.subscribers[eventType] = append(b.subscribers[eventType], listener)
	Log.WithFields(logrus.Fields{"module": listener.Module, "event": eventType}).Debug("Subscribed to event")

	return nil
}

// Subscribers returns the subscribers for a given event type.
func (b *EventBus) Subscribers(eventType EventType) []*Listener {
	return b.subscribers[eventType]
}

// Subscriber returns the subscriber for a given event type and listener ID.
func (b *EventBus) Subscriber(eventType EventType, listenerID string) *Listener {
	for _, l := range b.subscribers[eventType] {
		if l.ID == listenerID {
			return l
		}
	}
	return nil
}

// Unsubscribe unsubscribes a listener from a given event type.
func (b *EventBus) Unsubscribe(eventType EventType, listener Listener) error {

	// check if event type exists
	if _, ok := b.subscribers[eventType]; !ok {
		Log.WithFields(logrus.Fields{"module": listener.Module, "event": eventType}).Warn("event type not found")
		return errors.New("event type not found")
	}

	// remove listener from subscribers
	for i, l := range b.subscribers[eventType] {
		if l.ID == listener.ID {
			b.subscribers[eventType] = append(b.subscribers[eventType][:i], b.subscribers[eventType][i+1:]...)
			Log.WithFields(logrus.Fields{"module": listener.Module, "event": eventType}).Debug("Unsubscribed from event")
			return nil
		}
	}

	// listener not found
	Log.WithFields(logrus.Fields{"module": listener.Module, "event": eventType}).Warn("listener not found")
	return errors.New("listener not found")

}

// Publish publishes an event to the event bus.
func (b *EventBus) Publish(e Event) error {

	// check if event type exists
	if _, ok := b.subscribers[e.Type]; !ok {
		Log.WithFields(logrus.Fields{"event": e.Type}).Warn("event type not found")
		return errors.New("event type not found")
	}

	// publish event to subscribers
	for _, l := range b.subscribers[e.Type] {
		l.Handler(e)
	}

	return nil
}

func isValidListener(listener *Listener) error {
	switch {
	case listener == nil:
		return errors.New("listener is nil")
	case listener.Handler == nil:
		return errors.New("listener handler is nil")
	case listener.ID == "":
		return errors.New("listener ID is empty")
	case listener.Module == "":
		return errors.New("listener module is empty")
	default:
		return nil
	}
}
