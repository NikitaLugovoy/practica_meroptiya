// handler/event_image_handler.go
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

func PostNewEventImage(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var new_image model.EventImage

			err := json.NewDecoder(r.Body).Decode(&new_image)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			ids, err := repository.InsertNewEventImageInBD(pool, &new_image)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]int{"id": ids})

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func GetFullEventImages(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			images, err := repository.SelectFullEventImages(pool)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(images)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func GetEventImagesByEventId(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /event-images/<event_id> in handler", http.StatusBadRequest)
				return
			}
			event_id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			images, err := repository.SelectEventImagesByEventId(pool, event_id)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(images)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func UpdateEventImage(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /event-image/<id> in handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var new_image model.EventImage
			if err := json.NewDecoder(r.Body).Decode(&new_image); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			cnt := repository.UpdateOneEventImage(pool, id, &new_image)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cnt)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func DeleteEventImage(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /event-image/<id> in handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			cnt, _ := repository.DeleteOneEventImage(pool, id)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cnt)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}
