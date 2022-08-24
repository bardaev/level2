package main

import (
	"log"
	"net/http"
)

func main() {

	var storage Storage = NewMemoryStorage()
	var handler Handler = *NewHandler(storage)

	http.HandleFunc("/create_event", middleware(handler.createEvent))
	http.HandleFunc("/update_event", middleware(handler.updateEvent))
	http.HandleFunc("/delete_event", middleware(handler.deleteEvent))
	http.HandleFunc("/events_for_day", handler.eventsForDay)
	http.HandleFunc("/events_for_week", handler.eventsForWeek)
	http.HandleFunc("/events_for_monts", handler.eventsForMonth)

	http.ListenAndServe(":8080", nil)
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s, Endpoint: %s", r.Method, r.RequestURI)
		next(w, r)
	}
}
