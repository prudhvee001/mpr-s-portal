package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Notification struct {
	Message string `json:"message"`
}

var notifications []Notification

func getNotifications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}

func addNotification(w http.ResponseWriter, r *http.Request) {
	var n Notification
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	notifications = append(notifications, n)
	json.NewEncoder(w).Encode(notifications)
}

func main() {
	http.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getNotifications(w, r)
		} else if r.Method == http.MethodPost {
			addNotification(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("🚀 Notification service running on port 3003")
	log.Fatal(http.ListenAndServe(":3003", nil))
}
