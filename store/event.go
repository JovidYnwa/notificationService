package store

import (
	"sync"
	"time"
)

type Event struct {
	OrderType  string    `json:"orderType"`
	SessionID  string    `json:"sessionId"`
	Card       string    `json:"card"`
	EventDate  time.Time `json:"eventDate"`
	WebsiteURL string    `json:"websiteUrl"`
}

type EventStore interface {
	StoreEvent(event Event) error
	GetUnprocessedEvents() ([]Event, error)
}

type MemoryStore struct {
	mu     sync.RWMutex
	events []Event
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		events: make([]Event, 0),
	}
}

func (s *MemoryStore) StoreEvent(event Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
	return nil
}

func (s *MemoryStore) GetUnprocessedEvents() ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	events := s.events
	s.events = make([]Event, 0)
	return events, nil
}
