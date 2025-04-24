package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Notification struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func notificationsHandler(w http.ResponseWriter, r *http.Request) {
	notifications := []Notification{
		{ID: 1, Message: "Your marks have been published."},
		{ID: 2, Message: "Your attendance is 85% this month."},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}

func main() {
	http.HandleFunc("/notifications", notificationsHandler)
	log.Println("🚀 Notification service running on port 3003")
	log.Fatal(http.ListenAndServe(":3003", nil))
}
