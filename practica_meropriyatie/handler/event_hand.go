// handler/event_handler.go
package handler

import (
	"encoding/json"
	"net/http"
	"practica_meropriyatie/bd/model"
	"practica_meropriyatie/bd/repository"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func PostNewEvent(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var new_event model.Event

			err := json.NewDecoder(r.Body).Decode(&new_event)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			ids, err := repository.InsertNewEventInBD(pool, &new_event)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			user, err := repository.SelectOneUser(pool, int(new_event.GetResponsible_id()))
			if err != nil {
				http.Error(w, "user not found", http.StatusBadRequest)
				return
			}
			email := user.GetEmail()

			// отправка письма
			go email_send(email, new_event.GetName(), new_event.GetDate_time(), string(new_event.GetCategory()))

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]int{"id": ids})

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func GetFullEvents(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			events, err := repository.SelectFullEvents(pool)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(events)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func GetOneEvent(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /event/<id> in event handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			event, err := repository.SelectOneEvent(pool, id)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(event)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func UpdateEvent(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /event/<id> in event handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var new_event model.Event
			if err := json.NewDecoder(r.Body).Decode(&new_event); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			cnt := repository.UpdateOneEvent(pool, id, &new_event)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cnt)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func DeleteEvent(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /event/<id> in event handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			cnt, _ := repository.DeleteOneEvent(pool, id)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cnt)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func UpdateEventStatus(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /event-status/<id> in handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			cnt := repository.UpdateOneEventStatus_Done(pool, id)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cnt)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}
