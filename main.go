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

func HandleEvent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func main() {
	port := flag.String("port", ":3000", "server port")
	flag.Parse()

	http.HandleFunc("/event", HandleEvent)

	log.Printf("Server starting on %s...", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal(err)
	}
}
