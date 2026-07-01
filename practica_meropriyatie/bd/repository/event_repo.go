// repository/event_repository.go
package repository

import (
	"context"
	"fmt"
	"practica_meropriyatie/bd/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertNewEventInBD(pool *pgxpool.Pool, event *model.Event) (int, error) {
	rowInsert := pool.QueryRow(context.Background(),
		"INSERT INTO events(name, description, date_time, location, category_events, status, organizer_id, responsible_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id",
		event.GetName(),
		event.GetDescription(),
		event.GetDate_time(),
		event.GetLocation(),
		event.GetCategory(),
		event.GetStatus(),
		event.GetOrganizer_id(),
		event.GetResponsible_id(),
	)
	var idInsert int
	if err := rowInsert.Scan(&idInsert); err != nil {
		return 0, err
	}
	return idInsert, nil
}

func SelectFullEvents(pool *pgxpool.Pool) ([]model.Event, error) {

	var name, description, location, organizer_name, responsible_name string
	var category model.CategoryEvents
	var status model.EventStatus
	var id, organizer_id, responsible_id int
	var date_time, created_at time.Time

	rows, err := pool.Query(context.Background(),
		"SELECT events.id, events.name, events.description, events.date_time, events.location, events.category_events, events.status,"+
			" org.id, concat(org.name,' ',org.surname),"+
			" resp.id, concat(resp.name,' ',resp.surname),"+
			" events.created_at"+
			" FROM events"+
			" JOIN users org ON events.organizer_id = org.id"+
			" JOIN users resp ON events.responsible_id = resp.id")

	var events []model.Event
	for rows.Next() {
		if err := rows.Scan(&id, &name, &description, &date_time, &location, &category, &status,
			&organizer_id, &organizer_name,
			&responsible_id, &responsible_name,
			&created_at); err != nil {
			return nil, err
		}
		event := model.NewEvent(uint(id), name, description, date_time, location, category, status,
			organizer_id, organizer_name,
			responsible_id, responsible_name,
			created_at)
		events = append(events, *event)
	}
	rows.Close()
	fmt.Println(err)
	return events, nil
}

func SelectOneEvent(pool *pgxpool.Pool, ids int) (*model.Event, error) {

	var name, description, location, organizer_name, responsible_name string
	var category model.CategoryEvents
	var status model.EventStatus
	var id, organizer_id, responsible_id int
	var date_time, created_at time.Time

	row := pool.QueryRow(context.Background(),
		"SELECT events.id, events.name, events.description, events.date_time, events.location, events.category_events, events.status,"+
			" org.id, concat(org.name,' ',org.surname),"+
			" resp.id, concat(resp.name,' ',resp.surname),"+
			" events.created_at"+
			" FROM events"+
			" JOIN users org ON events.organizer_id = org.id"+
			" JOIN users resp ON events.responsible_id = resp.id"+
			" WHERE events.id = $1", ids)

	if err := row.Scan(&id, &name, &description, &date_time, &location, &category, &status,
		&organizer_id, &organizer_name,
		&responsible_id, &responsible_name,
		&created_at); err != nil {
		return nil, err
	}
	event := model.NewEvent(uint(id), name, description, date_time, location, category, status,
		organizer_id, organizer_name,
		responsible_id, responsible_name,
		created_at)
	return event, nil
}

func UpdateOneEvent(pool *pgxpool.Pool, ids int, event *model.Event) int64 {

	result, err := pool.Exec(context.Background(),
		"UPDATE events SET name=$1, description=$2, date_time=$3, location=$4, category_events=$5, status=$6, organizer_id=$7, responsible_id=$8"+
			" WHERE id = $9",
		event.GetName(),
		event.GetDescription(),
		event.GetDate_time(),
		event.GetLocation(),
		event.GetCategory(),
		event.GetStatus(),
		event.GetOrganizer_id(),
		event.GetResponsible_id(),
		ids)
	if err != nil {
		return 0
	}
	return result.RowsAffected()
}

// Изменение участника
func UpdateOneEventStatus_Done(pool *pgxpool.Pool, ids int) int64 {

	// Получаем текущий статус
	var current_status model.EventStatus
	row := pool.QueryRow(context.Background(),
		"SELECT status FROM events WHERE id = $1", ids)

	if err := row.Scan(&current_status); err != nil {
		return 0
	}

	// Переключаем статус
	var new_status model.EventStatus
	if current_status == model.StatusPlanned {
		new_status = model.StatusCompleted
	} else {
		new_status = model.StatusPlanned
	}

	// Обновляем
	result, err := pool.Exec(context.Background(),
		"UPDATE events SET status=$1 WHERE id = $2",
		new_status,
		ids)
	if err != nil {
		return 0
	}
	return result.RowsAffected()
}

func DeleteOneEvent(pool *pgxpool.Pool, ids int) (int64, error) {

	result, err := pool.Exec(context.Background(), "DELETE FROM events WHERE id = $1", ids)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), err
}
