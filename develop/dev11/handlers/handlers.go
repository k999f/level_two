package handlers

import (
	"bytes"
	"dev11/model"
	"dev11/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func CreateEvent(cache *[]model.Event) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод
		if r.Method != http.MethodPost {
			utils.ErrorResponse(w, 500, "Method not allowed")
			return
		}

		// Копируем request, чтобы повторно его читать
		body, _ := ioutil.ReadAll(r.Body)
		r2 := r.Clone(r.Context())
		r.Body = ioutil.NopCloser(bytes.NewReader(body))
		r2.Body = ioutil.NopCloser(bytes.NewReader(body))

		// Валидируем событие
		err, ok := utils.ValidateEvent(r2)
		if !ok {
			utils.ErrorResponse(w, 400, err)
			return
		}

		// Парсим событие
		event := utils.ParseEvent(r, len(*cache))

		// Добавляем событие в кэш
		*cache = append(*cache, event)

		// Отправляем ответ
		response := map[string]string{"result": "event created"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func DeleteEvent(cache *[]model.Event) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод
		if r.Method != http.MethodPost {
			utils.ErrorResponse(w, 500, "Method not allowed")
			return
		}

		eventData := model.EventDataId{}

		// Парсим событие
		err := json.NewDecoder(r.Body).Decode(&eventData)

		if err != nil {
			utils.ErrorResponse(w, 400, "Invalid JSON")
			return
		}

		// Ищем и удаляем событие в кэше
		tempCache := *cache
		found := false

		for i, event := range *cache {
			if event.EventId == eventData.EventId {
				*cache = append(tempCache[:i], tempCache[i+1:]...)
				found = true
				break
			}
		}

		if !found {
			utils.ErrorResponse(w, 404, "Event not found")
			return
		}

		// Отправляем ответ
		response := map[string]string{"result": "event deleted"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func UpdateEvent(cache *[]model.Event) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод
		if r.Method != http.MethodPost {
			utils.ErrorResponse(w, 500, "Method not allowed")
			return
		}

		// Копируем request, чтобы повторно его читать
		body, _ := ioutil.ReadAll(r.Body)
		r2 := r.Clone(r.Context())
		r.Body = ioutil.NopCloser(bytes.NewReader(body))
		r2.Body = ioutil.NopCloser(bytes.NewReader(body))

		// Валидируем событие
		err, ok := utils.ValidateUpdateEvent(r2)
		if !ok {
			utils.ErrorResponse(w, 400, err)
			return
		}

		// Парсим событие
		eventData := utils.ParseUpdateEvent(r, len(*cache))
		fmt.Println(eventData)

		// Ищем и изменяем событие в кэше
		tempCache := *cache
		found := false

		for i, event := range *cache {
			if event.EventId == eventData.EventId {
				tempCache[i].UserId = eventData.UserId
				tempCache[i].Description = eventData.Description
				tempCache[i].DateTime = eventData.DateTime
				*cache = tempCache
				found = true
				break
			}
		}

		if !found {
			utils.ErrorResponse(w, 404, "event not found")
			return
		}

		// Отправляем ответ
		response := map[string]string{"result": "event updated"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func EventsForDay(cache *[]model.Event) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод
		if r.Method != http.MethodGet {
			utils.ErrorResponse(w, 500, "Method not allowed")
			return
		}

		// Получаем дату
		qDate := r.URL.Query().Get("date")

		date, err := time.Parse("2006-01-02", qDate)

		if err != nil {
			utils.ErrorResponse(w, 400, "Invalid date")
			return
		}

		// Получаем пользователя
		qUserId := r.URL.Query().Get("user_id")

		userId, err := strconv.Atoi(qUserId)

		if err != nil {
			utils.ErrorResponse(w, 400, "Invalid user ID")
			return
		}

		var result []model.Event

		// Выбираем нужные события
		for _, event := range *cache {
			if event.DateTime.Year() == date.Year() && event.DateTime.Month() == date.Month() && event.DateTime.Day() == date.Day() {
				if event.UserId == userId {
					result = append(result, event)
				}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func EventsForPeriod(cache *[]model.Event, period string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод
		if r.Method != http.MethodGet {
			utils.ErrorResponse(w, 500, "Method not allowed")
			return
		}

		// Получаем дату
		qDate := r.URL.Query().Get("date")

		date, err := time.Parse("2006-01-02", qDate)

		if err != nil {
			utils.ErrorResponse(w, 400, "Invalid date")
			return
		}

		// Получаем пользователя
		qUserId := r.URL.Query().Get("user_id")

		userId, err := strconv.Atoi(qUserId)

		if err != nil {
			utils.ErrorResponse(w, 400, "Invalid user ID")
			return
		}

		// Выбираем диапазон поиска событий
		var start time.Time
		var end time.Time

		if period == "week" {
			start = date.AddDate(0, 0, -int(date.Weekday()))
			end = start.AddDate(0, 0, 6)
		} else if period == "month" {
			start = time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
			end = start.AddDate(0, 1, -1)
		}

		var result []model.Event

		// Выбираем нужные события
		for _, event := range *cache {
			if event.DateTime.After(start) && event.DateTime.Before(end) {
				if event.UserId == userId {
					result = append(result, event)
				}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

// func EventsForMonth(cache *[]model.Event) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Проверяем метод
// 		if r.Method != http.MethodGet {
// 			utils.ErrorResponse(w, 500, "Method not allowed")
// 			return
// 		}

// 		// Получаем дату
// 		qDate := r.URL.Query().Get("date")

// 		date, err := time.Parse("2006-01-02", qDate)

// 		if err != nil {
// 			utils.ErrorResponse(w, 400, "Invalid date")
// 			return
// 		}

// 		// Получаем пользователя
// 		qUserId := r.URL.Query().Get("user_id")

// 		userId, err := strconv.Atoi(qUserId)

// 		if err != nil {
// 			utils.ErrorResponse(w, 400, "Invalid user ID")
// 			return
// 		}

// 		start := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
// 		end := start.AddDate(0, 1, -1)

// 		var result []model.Event

// 		// Выбираем нужные события
// 		for _, event := range *cache {
// 			if event.DateTime.After(start) && event.DateTime.Before(end) {
// 				if event.UserId == userId {
// 					result = append(result, event)
// 				}
// 			}
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(result)
// 	}
// }
