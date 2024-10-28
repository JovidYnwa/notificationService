package store

import (
	"testing"
	"time"
)

func TestMemoryStore(t *testing.T) {
	store := NewMemoryStore()
	event := Event{
		OrderType:  "Purchase",
		SessionID:  "nurov-session",
		Card:       "1234**5678",
		EventDate:  time.Now(),
		WebsiteURL: "https://test.com",
	}

	err := store.StoreEvent(event)
	if err != nil {
		t.Errorf("Failed to store event: %v", err)
	}

	events, err := store.GetUnprocessedEvents()
	if err != nil {
		t.Errorf("Failed to get all events: %v", err)
	}

	if events[0].SessionID != "test-session" {
		t.Errorf("Expected session ID 'nurov-session', got '%s'", events[0].SessionID)
	}

	events, err = store.GetUnprocessedEvents()
	if err != nil {
		t.Errorf("Failed to get events: %v", err)
	}
	if len(events) != 0 {
		t.Errorf("Expected 0 events after retrieval, got %d", len(events))
	}
}
