package main

import (
	"fmt"
	"time"
)

type Event struct {
	OrderType  string    `json:"orderType"`
	SessionID  string    `json:"sessionId"`
	Card       string    `json:"card"`
	EventDate   time.Time `json:"eventDate"`
	WebsiteUrl string    `json:"websiteUrl"`
}

func main() {
	e1 := Event{
		OrderType:  "Purchase",
		SessionID:  "29827525-06c9-4b1e-9d9b-7c4584e82f56",
		Card:       "44331409",
		EventDate:  time.Now(),
		WebsiteUrl: "https://amazon.co",
	}
	fmt.Println(e1)
}
