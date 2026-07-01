package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practica_meropriyatie/bd/model"
	"practica_meropriyatie/bd/repository"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func UpdateUser(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /task/<id> in task handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var new_user model.User

			error := json.NewDecoder(r.Body).Decode(&new_user)
			if error != nil {
				http.Error(w, error.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			cnt := repository.UpdateOneUser(pool, id, &new_user)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(cnt)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func DeleteUser(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /task/<id> in task handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			cnt, _ := repository.DeleteOneUser(pool, id)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(cnt)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func PostNewUserRegister(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var new_user model.User

			err := json.NewDecoder(r.Body).Decode(&new_user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			full_user, _ := repository.SelectFullUser(pool)
			for _, user := range full_user {
				if user.Email == new_user.Email {
					http.Error(w, "ПОЛЬЗОВАТЕЛЬ ТАКОЙ УЖЕ СУЩЕСТВУЕТ!!!", http.StatusBadRequest)
					return
				}
			}

			ids, err := repository.InsertNewUserInBD(pool, &new_user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)

			json.NewEncoder(w).Encode(map[string]int{
				"id": ids,
			})
		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func LoginUser(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			var new_user model.User

			err := json.NewDecoder(r.Body).Decode(&new_user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			var check_user *model.User
			full_user, _ := repository.SelectFullUser(pool)
			fmt.Printf("Получен логин от клиента: %q\n", new_user.Login)
			for _, user := range full_user {
				fmt.Printf("Сравниваю с логином из БД: %q\n", user.Login)
				if user.Login == new_user.Login {
					check_user = &user
					break
				}
			}
			if check_user == nil {
				http.Error(w, "Такого пользователя не существует", http.StatusUnauthorized)
				return
			}

			if repository.HashPassword(new_user.Password) != check_user.Password {
				http.Error(w, "Неправильный пароль", http.StatusUnauthorized)
				return
			}

			token, err := GenerateJWT(int(check_user.GetId()))
			if err != nil {
				http.Error(w, "Ошибка генерации токена", 500)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(map[string]string{
				"token": token,
			})

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func GetOneUser(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

			path := strings.Trim(r.URL.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 2 {
				http.Error(w, "expect /task/<id> in task handler", http.StatusBadRequest)
				return
			}
			id, err := strconv.Atoi(pathParts[1])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			user, err := repository.SelectOneUser(pool, id)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(user)

		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}

func GetUsersByRole(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
			return
		}

		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		if len(parts) < 2 {
			http.Error(w, "expect /users_by_role/<role>", http.StatusBadRequest)
			return
		}

		role := parts[1]

		users, err := repository.SelectUsersByRole(pool, role)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func GetFullUser(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			users, err := repository.SelectFullUser(pool)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(users)
		} else {
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	}
}
