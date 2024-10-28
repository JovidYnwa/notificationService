package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"notification-service/store"
)

type Server struct {
	storage store.EventStore
}

func NewServer(storage store.EventStore) *Server {
	return &Server{
		storage: storage,
	}
}

func (s *Server) HandleEvent(w http.ResponseWriter, r *http.Request) {
	var event store.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := s.storage.StoreEvent(event); err != nil {
		http.Error(w, "Error storing event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func main() {
	port := flag.String("port", ":3000", "server port")
	flag.Parse()

	storage := store.NewMemoryStore()

	server := NewServer(storage)
	http.HandleFunc("/event", server.HandleEvent)

	log.Printf("Server starting on %s", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal(err)
	}
}
