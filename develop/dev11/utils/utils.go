package utils

import (
	"dev11/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func ErrorResponse(w http.ResponseWriter, status int, error string) {
	response := map[string]string{"error": error}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func ValidateEvent(r *http.Request) (string, bool) {
	eventData := model.EventData{}

	err := json.NewDecoder(r.Body).Decode(&eventData)

	if err != nil {
		return "Invalid JSON", false
	}

	_, err = time.Parse("2006-01-02", eventData.DateTime)

	if err != nil {
		return "Invalid date", false
	}

	return "", true
}

func ValidateUpdateEvent(r *http.Request) (string, bool) {
	eventData := model.EventDataUpdate{}

	err := json.NewDecoder(r.Body).Decode(&eventData)

	if err != nil {
		return "Invalid JSON", false
	}

	_, err = time.Parse("2006-01-02", eventData.DateTime)

	if err != nil {
		return "Invalid date", false
	}

	return "", true
}

func ParseEvent(r *http.Request, cacheLen int) model.Event {
	eventData := model.EventData{}

	json.NewDecoder(r.Body).Decode(&eventData)

	date, _ := time.Parse("2006-01-02", eventData.DateTime)

	event := model.Event{
		EventId:     cacheLen + 1,
		UserId:      eventData.UserId,
		DateTime:    date,
		Description: eventData.Description,
	}

	fmt.Println("In func: ", event)

	return event
}

func ParseUpdateEvent(r *http.Request, cacheLen int) model.Event {
	eventData := model.EventDataUpdate{}

	json.NewDecoder(r.Body).Decode(&eventData)

	date, _ := time.Parse("2006-01-02", eventData.DateTime)

	event := model.Event{
		EventId:     eventData.EventId,
		UserId:      eventData.UserId,
		DateTime:    date,
		Description: eventData.Description,
	}

	fmt.Println("In func: ", event)

	return event
}

func MiddlewareLogger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("URL: %s; method: %s\n", r.URL.RequestURI(), r.Method)
		handler(w, r)
	}
}
