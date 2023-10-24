package core

import "testing"

func TestSubscribeAndPublish(t *testing.T) {

	tests := []struct {
		name      string
		eventType EventType
		want      bool
	}{
		{"EventRecieved", "testEvent", true},
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
			bus.Publish(Event{Type: tt.eventType})

			if recieved != tt.want {
				t.Errorf("Subscribe() = %v, want %v", recieved, tt.want)
			}
		})
	}
}

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
