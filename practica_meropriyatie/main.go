package main

import (
	"fmt"
	"log"
	"net/http"
	db "practica_meropriyatie/bd"
	"practica_meropriyatie/handler"
)

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// preflight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func jwt(next http.HandlerFunc) http.HandlerFunc {
	return enableCORS(handler.JWTMiddleware(next))
}

func main() {

	con_postgres := db.AddDataBase()
	defer con_postgres.Close()
	db.RunMigration(con_postgres)

	// ===== USERS =====
	http.HandleFunc("/full_user", jwt(handler.GetFullUser(con_postgres)))
	http.HandleFunc("/users_by_role/", jwt(handler.GetUsersByRole(con_postgres)))
	http.HandleFunc("/new_user", enableCORS(handler.PostNewUserRegister(con_postgres))) // публичный — регистрация
	http.HandleFunc("/login", enableCORS(handler.LoginUser(con_postgres)))              // публичный — логин
	http.HandleFunc("/one_user/", jwt(handler.GetOneUser(con_postgres)))
	http.HandleFunc("/delete_user/", jwt(handler.DeleteUser(con_postgres)))
	http.HandleFunc("/update_user/", jwt(handler.UpdateUser(con_postgres)))

	// ===== EVENTS =====
	http.HandleFunc("/full_event", jwt(handler.GetFullEvents(con_postgres)))
	http.HandleFunc("/one_event/", jwt(handler.GetOneEvent(con_postgres)))
	http.HandleFunc("/new_event", jwt(handler.PostNewEvent(con_postgres)))
	http.HandleFunc("/delete_event/", jwt(handler.DeleteEvent(con_postgres)))
	http.HandleFunc("/update_event/", jwt(handler.UpdateEvent(con_postgres)))
	http.HandleFunc("/event_status/", jwt(handler.UpdateEventStatus(con_postgres)))

	// ===== EVENT IMAGES =====
	http.HandleFunc("/full_event_image", jwt(handler.GetFullEventImages(con_postgres)))
	http.HandleFunc("/event_image_by_event/", jwt(handler.GetEventImagesByEventId(con_postgres)))
	http.HandleFunc("/new_event_image", jwt(handler.PostNewEventImage(con_postgres)))
	http.HandleFunc("/update_event_image/", jwt(handler.UpdateEventImage(con_postgres)))
	http.HandleFunc("/delete_event_image/", jwt(handler.DeleteEventImage(con_postgres)))

	// ===== EVENT PARTICIPANTS =====
	http.HandleFunc("/full_event_participant", jwt(handler.GetFullEventParticipants(con_postgres)))
	http.HandleFunc("/one_event_participant/", jwt(handler.GetOneEventParticipant(con_postgres)))
	http.HandleFunc("/participants_by_event/", jwt(handler.GetParticipantsByEventId(con_postgres)))
	http.HandleFunc("/new_event_participant", jwt(handler.PostNewEventParticipant(con_postgres)))
	http.HandleFunc("/add_event_participants_by_group", jwt(handler.AddParticipantsByGroup(con_postgres)))
	http.HandleFunc("/update_event_participant/", jwt(handler.UpdateEventParticipant(con_postgres)))
	http.HandleFunc("/toggle_participant_status/", jwt(handler.UpdateEventParticipantStatus(con_postgres)))
	http.HandleFunc("/delete_event_participant/", jwt(handler.DeleteEventParticipant(con_postgres)))

	// ===== GROUPS =====
	http.HandleFunc("/full_group", jwt(handler.GetFullGroups(con_postgres)))
	http.HandleFunc("/one_group/", jwt(handler.GetOneGroup(con_postgres)))
	http.HandleFunc("/new_group", jwt(handler.PostNewGroup(con_postgres)))
	http.HandleFunc("/update_group/", jwt(handler.UpdateGroup(con_postgres)))
	http.HandleFunc("/delete_group/", jwt(handler.DeleteGroup(con_postgres)))

	// ===== USER GROUPS =====
	http.HandleFunc("/full_user_group", jwt(handler.GetFullUserGroups(con_postgres)))
	http.HandleFunc("/one_user_group/", jwt(handler.GetOneUserGroup(con_postgres)))
	http.HandleFunc("/user_groups_by_group/", jwt(handler.GetUserGroupsByGroupId(con_postgres)))
	http.HandleFunc("/new_user_group", jwt(handler.PostNewUserGroup(con_postgres)))
	http.HandleFunc("/update_user_group/", jwt(handler.UpdateUserGroup(con_postgres)))
	http.HandleFunc("/delete_user_group/", jwt(handler.DeleteUserGroup(con_postgres)))

	// ===== EVENT RESULTS =====
	http.HandleFunc("/full_event_result", jwt(handler.GetFullEventResults(con_postgres)))
	http.HandleFunc("/one_event_result/", jwt(handler.GetOneEventResult(con_postgres)))
	http.HandleFunc("/event_results_by_event/", jwt(handler.GetEventResultsByEventId(con_postgres)))
	http.HandleFunc("/new_event_result", jwt(handler.PostNewEventResult(con_postgres)))
	http.HandleFunc("/update_event_result/", jwt(handler.UpdateEventResult(con_postgres)))
	http.HandleFunc("/delete_event_result/", jwt(handler.DeleteEventResult(con_postgres)))

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
