package model

import "time"

type Event struct {
	EventId     int       `json:"event_id"`
	UserId      int       `json:"user_id"`
	DateTime    time.Time `json:"date_time"`
	Description string    `json:"description"`
}

type EventData struct {
	UserId      int    `json:"user_id"`
	DateTime    string `json:"date_time"`
	Description string `json:"description"`
}

type EventDataId struct {
	EventId int `json:"event_id"`
}

type EventDataUpdate struct {
	EventId     int    `json:"event_id"`
	UserId      int    `json:"user_id"`
	DateTime    string `json:"date_time"`
	Description string `json:"description"`
}
