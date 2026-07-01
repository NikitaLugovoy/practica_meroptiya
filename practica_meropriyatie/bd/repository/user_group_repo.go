package repository

import (
	"context"
	"fmt"
	"practica_meropriyatie/bd/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertNewUserGroupInBD(pool *pgxpool.Pool, user_group *model.UserGroup) (int, error) {
	rowInsert := pool.QueryRow(context.Background(),
		"INSERT INTO user_groups(user_id, group_id) VALUES ($1,$2) RETURNING id",
		user_group.GetUser_id(),
		user_group.GetGroup_id(),
	)
	var idInsert int
	if err := rowInsert.Scan(&idInsert); err != nil {
		return 0, err
	}
	return idInsert, nil
}

func SelectFullUserGroups(pool *pgxpool.Pool) ([]model.UserGroup, error) {

	var user_name, group_name string
	var id, user_id, group_id int
	var created_at time.Time

	rows, err := pool.Query(context.Background(),
		"SELECT ug.id, users.id, concat(users.name,' ',users.surname), groups.id, groups.name, ug.created_at"+
			" FROM user_groups ug"+
			" JOIN users ON ug.user_id = users.id"+
			" JOIN groups ON ug.group_id = groups.id")

	var user_groups []model.UserGroup
	for rows.Next() {
		if err := rows.Scan(&id, &user_id, &user_name, &group_id, &group_name, &created_at); err != nil {
			return nil, err
		}
		user_group := model.NewUserGroup(uint(id), user_id, user_name, group_id, group_name, created_at)
		user_groups = append(user_groups, *user_group)
	}
	rows.Close()
	fmt.Println(err)
	return user_groups, nil
}

func SelectOneUserGroup(pool *pgxpool.Pool, ids int) (*model.UserGroup, error) {

	var user_name, group_name string
	var id, user_id, group_id int
	var created_at time.Time

	row := pool.QueryRow(context.Background(),
		"SELECT ug.id, users.id, concat(users.name,' ',users.surname), groups.id, groups.name, ug.created_at"+
			" FROM user_groups ug"+
			" JOIN users ON ug.user_id = users.id"+
			" JOIN groups ON ug.group_id = groups.id"+
			" WHERE ug.id = $1", ids)

	if err := row.Scan(&id, &user_id, &user_name, &group_id, &group_name, &created_at); err != nil {
		return nil, err
	}
	user_group := model.NewUserGroup(uint(id), user_id, user_name, group_id, group_name, created_at)
	return user_group, nil
}

// Вывод всех записей по id группы
func SelectUserGroupsByGroupId(pool *pgxpool.Pool, group_id int) ([]model.UserGroup, error) {

	var user_name, group_name string
	var id, user_id int
	var created_at time.Time

	rows, err := pool.Query(context.Background(),
		"SELECT ug.id, users.id, concat(users.name,' ',users.surname), groups.id, groups.name, ug.created_at"+
			" FROM user_groups ug"+
			" JOIN users ON ug.user_id = users.id"+
			" JOIN groups ON ug.group_id = groups.id"+
			" WHERE ug.group_id = $1", group_id)

	var user_groups []model.UserGroup
	for rows.Next() {
		if err := rows.Scan(&id, &user_id, &user_name, &group_id, &group_name, &created_at); err != nil {
			return nil, err
		}
		user_group := model.NewUserGroup(uint(id), user_id, user_name, group_id, group_name, created_at)
		user_groups = append(user_groups, *user_group)
	}
	rows.Close()
	fmt.Println(err)
	return user_groups, nil
}

// Изменение записи
func UpdateOneUserGroup(pool *pgxpool.Pool, ids int, user_group *model.UserGroup) int64 {

	result, err := pool.Exec(context.Background(),
		"UPDATE user_groups SET user_id=$1, group_id=$2 WHERE id = $3",
		user_group.GetUser_id(),
		user_group.GetGroup_id(),
		ids)
	if err != nil {
		return 0
	}
	return result.RowsAffected()
}

func DeleteOneUserGroup(pool *pgxpool.Pool, ids int) (int64, error) {

	result, err := pool.Exec(context.Background(), "DELETE FROM user_groups WHERE id = $1", ids)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), err
}
