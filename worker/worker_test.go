package worker

import (
	"notification-service/store"
	"testing"
	"time"
)

type MockStorage struct {
	events []store.Event
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		events: make([]store.Event, 0),
	}
}

func (m *MockStorage) StoreEvent(event store.Event) error {
	m.events = append(m.events, event)
	return nil
}

func (m *MockStorage) GetUnprocessedEvents() ([]store.Event, error) {
	events := m.events
	m.events = make([]store.Event, 0)
	return events, nil
}

func TestWorker(t *testing.T) {
	mockStorage := NewMockStorage()
	worker := NewNotificationWorker(mockStorage, 100*time.Millisecond)

	testEvent := store.Event{
		OrderType:  "Purchase",
		SessionID:  "test-session",
		Card:       "1234**5678",
		EventDate:  time.Now(),
		WebsiteURL: "https://test.com",
	}

	mockStorage.StoreEvent(testEvent)

	go worker.Start()

	time.Sleep(200 * time.Millisecond)

	worker.Stop()

	events, _ := mockStorage.GetUnprocessedEvents()
	if len(events) != 0 {
		t.Errorf("Expected 0 events after processing, got %d", len(events))
	}
}
