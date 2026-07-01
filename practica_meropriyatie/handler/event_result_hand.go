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

func PostNewEventResult(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var new_result model.EventResult

			err := json.NewDecoder(r.Body).Decode(&new_result)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			ids, err := repository.InsertNewEventResultInBD(pool, &new_result)
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

func GetFullEventResults(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			results, err := repository.SelectFullEventResults(pool)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(results)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func GetOneEventResult(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /one_event_result/<id> in handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			result, err := repository.SelectOneEventResult(pool, id)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func GetEventResultsByEventId(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /event_results_by_event/<event_id> in handler", http.StatusBadRequest)
				return
			}
			event_id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			results, err := repository.SelectEventResultsByEventId(pool, event_id)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(results)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func UpdateEventResult(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /update_event_result/<id> in handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var new_result model.EventResult
			if err := json.NewDecoder(r.Body).Decode(&new_result); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			cnt := repository.UpdateOneEventResult(pool, id, &new_result)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cnt)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func DeleteEventResult(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /delete_event_result/<id> in handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			cnt, _ := repository.DeleteOneEventResult(pool, id)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cnt)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}
