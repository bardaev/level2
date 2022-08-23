package main

import "net/http"

func main() {

	var storage Storage = NewMemoryStorage()
	var handler Handler = *NewHandler(storage)

	http.HandleFunc("/create_event", handler.createEvent)
	http.HandleFunc("/update_event", handler.updateEvent)
	http.HandleFunc("/delete_event", handler.deleteEvent)
	http.HandleFunc("/events_for_day", handler.eventsForDay)
	http.HandleFunc("/events_for_week", handler.eventsForWeek)
	http.HandleFunc("/events_for_monts", handler.eventsForMonth)

	http.ListenAndServe(":8080", nil)
}
