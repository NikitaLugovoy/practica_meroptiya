// handler/event_participant_handler.go
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

func PostNewEventParticipant(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var new_participant model.EventParticipant

			err := json.NewDecoder(r.Body).Decode(&new_participant)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			role, errors := repository.CheckUserRole(pool, new_participant.User_id)
			if errors != nil {
				http.Error(w, "Пользователь не найден", http.StatusBadRequest)
				return
			}

			if role != "student" {
				http.Error(w,
					"Пользователь должен иметь роль student",
					http.StatusBadRequest)
				return
			}

			user, err := repository.SelectOneUser(pool, int(new_participant.User_id))
			if err != nil {
				http.Error(w, "user not found", http.StatusBadRequest)
				return
			}
			email := user.GetEmail()
			new_event, err := repository.SelectOneEvent(pool, int(new_participant.Event_id))
			if err != nil {
				http.Error(w, "event not found", http.StatusBadRequest)
				return
			}
			go email_send(email, new_event.GetName(), new_event.GetDate_time(), string(new_event.GetCategory()))

			ids, err := repository.InsertNewEventParticipantInBD(pool, &new_participant)
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

func AddParticipantsByGroup(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			EventID int `json:"event_id"`
			GroupID int `json:"group_id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if req.EventID == 0 || req.GroupID == 0 {
			http.Error(w, "event_id and group_id are required", http.StatusBadRequest)
			return
		}

		cnt, err := repository.InsertEventParticipantsByGroup(pool, req.EventID, req.GroupID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(map[string]int64{
			"inserted": cnt,
		})
	}
}

func GetFullEventParticipants(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			participants, err := repository.SelectFullEventParticipants(pool)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(participants)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func GetOneEventParticipant(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /event-participant/<id> in handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			participant, err := repository.SelectOneEventParticipant(pool, id)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(participant)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func GetParticipantsByEventId(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /event-participants/<event_id> in handler", http.StatusBadRequest)
				return
			}
			event_id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			participants, err := repository.SelectParticipantsByEventId(pool, event_id)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(participants)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func UpdateEventParticipant(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /event-participant/<id> in handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var new_participant model.EventParticipant
			if err := json.NewDecoder(r.Body).Decode(&new_participant); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			cnt := repository.UpdateOneEventParticipant(pool, id, &new_participant)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cnt)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func UpdateEventParticipantStatus(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /event-participant-status/<id> in handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			cnt := repository.UpdateOneEventStatus_OKe(pool, id)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cnt)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func DeleteEventParticipant(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /event-participant/<id> in handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			cnt, _ := repository.DeleteOneEventParticipant(pool, id)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cnt)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}
