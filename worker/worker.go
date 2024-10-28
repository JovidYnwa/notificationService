package worker

import (
	"log"
	"notification-service/store"
	"time"
)

type NotificationWorker struct {
	storage store.EventStore
	ticker  *time.Ticker
}

func NewNotificationWorker(storage store.EventStore, interval time.Duration) *NotificationWorker {
	return &NotificationWorker{
		storage: storage,
		ticker:  time.NewTicker(interval),
	}
}

func (w *NotificationWorker) Start() {
	log.Println("Worker up...")
	for range w.ticker.C {
		events, err := w.storage.GetUnprocessedEvents()
		if err != nil {
			log.Printf("Error getting events: %v", err)
			continue
		}
		for _, event := range events {
			w.processEvent(event)
		}
	}
}

func (w *NotificationWorker) Stop() {
	w.ticker.Stop()
}

func (w *NotificationWorker) processEvent(event store.Event) {
	log.Printf("\nNOTIFICATION:\n"+
		"Card %s was used for %s\n"+
		"Website: %s\n"+
		"Date: %s\n"+
		"Session: %s\n",
		event.Card,
		event.OrderType,
		event.WebsiteURL,
		event.EventDate.Format(time.RFC3339),
		event.SessionID)
}
