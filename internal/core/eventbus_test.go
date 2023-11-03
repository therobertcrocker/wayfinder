package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSubscribe tests the Subscribe function.
func TestSubscribe(t *testing.T) {
	tests := []struct {
		name      string
		eventType EventType
		expectErr bool
	}{
		{"SubscribeToEventType", "testEvent", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			bus := NewEventBus()

			listener := &Listener{
				ID:      "test",
				Module:  "test",
				Handler: func(e Event) {},
			}

			err := bus.Subscribe(tt.eventType, listener)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				sub := bus.Subscriber(tt.eventType, listener.ID)

				assert.NotNil(t, sub)
				assert.Equal(t, listener, sub)
			}

		})

	}

}

// TestPublish tests the Publish function.
func TestPublish(t *testing.T) {

	tests := []struct {
		name         string
		eventType    EventType
		wantResponse bool
		expectErr    bool
	}{
		{"EventRecieved", "testEvent", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := NewEventBus()
			recieved := false

			listener := &Listener{
				ID:      "test",
				Module:  "test",
				Handler: func(e Event) { recieved = true },
			}

			bus.Subscribe(tt.eventType, listener)
			err := bus.Publish(Event{Type: tt.eventType})

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.wantResponse, recieved)
			}

		})
	}
}

// TestUnsubscribe tests the Unsubscribe function.
func TestUnsubscribe(t *testing.T) {

	tests := []struct {
		name      string
		eventType EventType
		want      bool
	}{
		{"EventRecieved", "testEvent", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bus := NewEventBus()
			recieved := false

			listener := &Listener{
				ID:      "test",
				Module:  "test",
				Handler: func(e Event) { recieved = true },
			}

			bus.Subscribe(tt.eventType, listener)
			bus.Unsubscribe(tt.eventType, *listener)
			bus.Publish(Event{Type: tt.eventType})

			if recieved != tt.want {
				t.Errorf("Subscribe() = %v, want %v", recieved, tt.want)
			}
		})
	}
}
