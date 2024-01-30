package main

import (
	"dev11/config"
	"dev11/handlers"
	"dev11/model"
	"dev11/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Читаем конфиг
	configPath := "config/config.json"
	config := config.InitialzeConfig(configPath)
	// Создаем кэш
	var cache []model.Event

	address := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)

	// Ставим обработчики
	http.HandleFunc("/create_event", utils.MiddlewareLogger(handlers.CreateEvent(&cache)))
	http.HandleFunc("/delete_event", utils.MiddlewareLogger(handlers.DeleteEvent(&cache)))
	http.HandleFunc("/update_event", utils.MiddlewareLogger(handlers.UpdateEvent(&cache)))
	http.HandleFunc("/events_for_day", utils.MiddlewareLogger(handlers.EventsForDay(&cache)))
	http.HandleFunc("/events_for_week", utils.MiddlewareLogger(handlers.EventsForPeriod(&cache, "week")))
	http.HandleFunc("/events_for_month", utils.MiddlewareLogger(handlers.EventsForPeriod(&cache, "month")))

	log.Printf("Server is running on %s", address)

	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal("Server listengng and serving error: ", err)
	}
}
