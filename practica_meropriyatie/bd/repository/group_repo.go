// repository/group_repo.go
package repository

import (
	"context"
	"fmt"
	"practica_meropriyatie/bd/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertNewGroupInBD(pool *pgxpool.Pool, group *model.Group) (int, error) {
	rowInsert := pool.QueryRow(context.Background(),
		"INSERT INTO groups(name) VALUES ($1) RETURNING id",
		group.GetName(),
	)
	var idInsert int
	if err := rowInsert.Scan(&idInsert); err != nil {
		return 0, err
	}
	return idInsert, nil
}

func SelectFullGroups(pool *pgxpool.Pool) ([]model.Group, error) {

	var name string
	var id int
	var created_at time.Time

	rows, err := pool.Query(context.Background(),
		"SELECT id, name, created_at FROM groups")

	var groups []model.Group
	for rows.Next() {
		if err := rows.Scan(&id, &name, &created_at); err != nil {
			return nil, err
		}
		group := model.NewGroup(uint(id), name, created_at)
		groups = append(groups, *group)
	}
	rows.Close()
	fmt.Println(err)
	return groups, nil
}

func SelectOneGroup(pool *pgxpool.Pool, ids int) (*model.Group, error) {

	var name string
	var id int
	var created_at time.Time

	row := pool.QueryRow(context.Background(),
		"SELECT id, name, created_at FROM groups WHERE id = $1", ids)

	if err := row.Scan(&id, &name, &created_at); err != nil {
		return nil, err
	}
	group := model.NewGroup(uint(id), name, created_at)
	return group, nil
}

func UpdateOneGroup(pool *pgxpool.Pool, ids int, group *model.Group) int64 {

	result, err := pool.Exec(context.Background(),
		"UPDATE groups SET name=$1 WHERE id = $2",
		group.GetName(),
		ids)
	if err != nil {
		return 0
	}
	return result.RowsAffected()
}

func DeleteOneGroup(pool *pgxpool.Pool, ids int) (int64, error) {

	result, err := pool.Exec(context.Background(), "DELETE FROM groups WHERE id = $1", ids)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), err
}
