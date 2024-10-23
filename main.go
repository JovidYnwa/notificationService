package main

import (
	"fmt"
	"net/http"
	"time"
)

type Event struct {
	OrderType  string    `json:"orderType"`
	SessionID  string    `json:"sessionId"`
	Card       string    `json:"card"`
	EventDate  time.Time `json:"eventDate"`
	WebsiteUrl string    `json:"websiteUrl"`
}

func YO(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "yo")
}

func main() {
	http.HandleFunc("/", Yo)
	fmt.Println("server is on port 3000!")
	http.ListenAndServe(":3000", nil)
}
