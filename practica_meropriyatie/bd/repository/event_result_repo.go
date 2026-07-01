package repository

import (
	"context"
	"fmt"
	"practica_meropriyatie/bd/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertNewEventResultInBD(pool *pgxpool.Pool, result *model.EventResult) (int, error) {
	rowInsert := pool.QueryRow(context.Background(),
		"INSERT INTO events_result(event_id, result) VALUES ($1,$2) RETURNING id",
		result.GetEvent_id(),
		result.GetResult(),
	)
	var idInsert int
	if err := rowInsert.Scan(&idInsert); err != nil {
		return 0, err
	}
	return idInsert, nil
}

func SelectFullEventResults(pool *pgxpool.Pool) ([]model.EventResult, error) {

	var result string
	var id, event_id int
	var created_at time.Time

	rows, err := pool.Query(context.Background(),
		"SELECT id, event_id, result, created_at FROM events_result")

	var results []model.EventResult
	for rows.Next() {
		if err := rows.Scan(&id, &event_id, &result, &created_at); err != nil {
			return nil, err
		}
		eventResult := model.NewEventResult(uint(id), event_id, result, created_at)
		results = append(results, *eventResult)
	}
	rows.Close()
	fmt.Println(err)
	return results, nil
}

func SelectEventResultsByEventId(pool *pgxpool.Pool, event_id int) ([]model.EventResult, error) {

	var result string
	var id int
	var created_at time.Time

	rows, err := pool.Query(context.Background(),
		"SELECT id, event_id, result, created_at FROM events_result WHERE event_id = $1", event_id)

	var results []model.EventResult
	for rows.Next() {
		if err := rows.Scan(&id, &event_id, &result, &created_at); err != nil {
			return nil, err
		}
		eventResult := model.NewEventResult(uint(id), event_id, result, created_at)
		results = append(results, *eventResult)
	}
	rows.Close()
	fmt.Println(err)
	return results, nil
}

func SelectOneEventResult(pool *pgxpool.Pool, ids int) (*model.EventResult, error) {

	var result string
	var id, event_id int
	var created_at time.Time

	row := pool.QueryRow(context.Background(),
		"SELECT id, event_id, result, created_at FROM events_result WHERE id = $1", ids)

	if err := row.Scan(&id, &event_id, &result, &created_at); err != nil {
		return nil, err
	}
	eventResult := model.NewEventResult(uint(id), event_id, result, created_at)
	return eventResult, nil
}

func UpdateOneEventResult(pool *pgxpool.Pool, ids int, result *model.EventResult) int64 {

	res, err := pool.Exec(context.Background(),
		"UPDATE events_result SET event_id=$1, result=$2 WHERE id = $3",
		result.GetEvent_id(),
		result.GetResult(),
		ids)
	if err != nil {
		return 0
	}
	return res.RowsAffected()
}

func DeleteOneEventResult(pool *pgxpool.Pool, ids int) (int64, error) {

	result, err := pool.Exec(context.Background(), "DELETE FROM events_result WHERE id = $1", ids)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), err
}
