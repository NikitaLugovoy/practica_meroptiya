package repository

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"practica_meropriyatie/bd/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func HashPassword(password string) string {
	byte_password := []byte(password)
	password_byte_to_hash := sha256.Sum256(byte_password)
	return hex.EncodeToString(password_byte_to_hash[:])
}

func InsertNewUserInBD(pool *pgxpool.Pool, user *model.User) (int, error) {
	rowInsert := pool.QueryRow(context.Background(),
		`INSERT INTO users(name, surname, login, password, phone_number, email, url_avatar, user_role) 
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`,
		user.GetName(),
		user.GetSurname(),
		user.GetLogin(),
		HashPassword(user.GetPassword()),
		user.GetPhoneNumber(),
		user.GetEmail(),
		user.GetUrl_avatar(),
		user.GetUser_role(),
	)

	var idInsert int
	if err := rowInsert.Scan(&idInsert); err != nil {
		return 0, err
	}
	return idInsert, nil
}

func SelectFullUser(pool *pgxpool.Pool) ([]model.User, error) {
	var (
		id                                                  int
		name, surname, login, password, phone_number, email string
		url_avatar, user_role                               string
		created_at                                          time.Time
	)

	rows, err := pool.Query(context.Background(),
		`SELECT id, name, surname, login, password, phone_number, email, url_avatar, user_role, created_at 
		 FROM users`)
	var users []model.User
	for rows.Next() {
		if err := rows.Scan(
			&id, &name, &surname, &login, &password,
			&phone_number, &email, &url_avatar, &user_role, &created_at,
		); err != nil {
			return nil, err
		}

		user := model.NewUser(uint(id), name, surname, login, password, phone_number, email, url_avatar, user_role, created_at)
		users = append(users, *user)
	}
	rows.Close()
	fmt.Println(err)
	return users, nil
}

func SelectOneUser(pool *pgxpool.Pool, ids int) (*model.User, error) {
	var (
		id                                                  int
		name, surname, login, password, phone_number, email string
		url_avatar, user_role                               string
		created_at                                          time.Time
	)

	rows := pool.QueryRow(context.Background(),
		`SELECT id, name, surname, login, password, phone_number, email, url_avatar, user_role, created_at 
		 FROM users WHERE id = $1`, ids)

	if err := rows.Scan(
		&id, &name, &surname, &login, &password,
		&phone_number, &email, &url_avatar, &user_role, &created_at,
	); err != nil {
		return nil, err
	}

	return model.NewUser(uint(id), name, surname, login, password, phone_number, email, url_avatar, user_role, created_at), nil
}

func SelectUsersByRole(pool *pgxpool.Pool, role string) ([]model.User, error) {
	var users []model.User

	rows, err := pool.Query(context.Background(),
		`SELECT id, name, surname, login, password, phone_number, email, url_avatar, user_role, created_at
		 FROM users WHERE user_role = $1`, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id                                                  int
			name, surname, login, password, phone_number, email string
			url_avatar, user_role                               string
			created_at                                          time.Time
		)

		if err := rows.Scan(
			&id, &name, &surname, &login, &password,
			&phone_number, &email, &url_avatar, &user_role, &created_at,
		); err != nil {
			return nil, err
		}

		users = append(users, *model.NewUser(uint(id), name, surname, login, password, phone_number, email, url_avatar, user_role, created_at))
	}

	return users, nil
}

func SelectUserByLogin(pool *pgxpool.Pool, login string) (*model.User, error) {
	var (
		id                                                    int
		name, surname, loginDB, password, phone_number, email string
		url_avatar, user_role                                 string
		created_at                                            time.Time
	)

	row := pool.QueryRow(context.Background(),
		`SELECT id, name, surname, login, password, phone_number, email, url_avatar, user_role, created_at 
		 FROM users WHERE login = $1`, login)

	if err := row.Scan(
		&id, &name, &surname, &loginDB, &password,
		&phone_number, &email, &url_avatar, &user_role, &created_at,
	); err != nil {
		return nil, err
	}

	return model.NewUser(uint(id), name, surname, loginDB, password, phone_number, email, url_avatar, user_role, created_at), nil
}

func CheckUserRole(pool *pgxpool.Pool, userID int) (string, error) {
	var role string
	err := pool.QueryRow(context.Background(),
		"SELECT user_role FROM users WHERE id=$1", userID,
	).Scan(&role)
	return role, err
}

func UpdateOneUser(pool *pgxpool.Pool, ids int, user *model.User) int64 {
	result, err := pool.Exec(context.Background(),
		`UPDATE users 
		 SET name=$1, surname=$2, login=$3, password=$4, phone_number=$5, email=$6, url_avatar=$7, user_role=$8
		 WHERE id = $9`,
		user.GetName(),
		user.GetSurname(),
		user.GetLogin(),
		HashPassword(user.GetPassword()),
		user.GetPhoneNumber(),
		user.GetEmail(),
		user.GetUrl_avatar(),
		user.GetUser_role(),
		ids,
	)
	if err != nil {
		return 0
	}

	return result.RowsAffected()
}

func DeleteOneUser(pool *pgxpool.Pool, ids int) (int64, error) {
	result, err := pool.Exec(context.Background(), "DELETE FROM users WHERE id = $1", ids)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), err
}
