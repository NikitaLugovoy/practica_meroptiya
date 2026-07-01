// repository/event_participant_repository.go
package repository

import (
	"context"
	"fmt"
	"practica_meropriyatie/bd/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertNewEventParticipantInBD(pool *pgxpool.Pool, participant *model.EventParticipant) (int, error) {
	rowInsert := pool.QueryRow(context.Background(),
		"INSERT INTO event_participants(event_id, user_id) VALUES ($1,$2) RETURNING id",
		participant.GetEvent_id(),
		participant.GetUser_id(),
	)
	var idInsert int
	if err := rowInsert.Scan(&idInsert); err != nil {
		return 0, err
	}
	return idInsert, nil
}

func InsertEventParticipantsByGroup(pool *pgxpool.Pool, event_id int, group_id int) (int64, error) {

	result, err := pool.Exec(context.Background(),
		`INSERT INTO event_participants (event_id, user_id)
		 SELECT $1, ug.user_id
		 FROM user_groups ug
		 WHERE ug.group_id = $2`,
		event_id,
		group_id,
	)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func SelectFullEventParticipants(pool *pgxpool.Pool) ([]model.EventParticipant, error) {

	var event_name, user_name string
	var participants_status model.ParticipantStatus
	var id, event_id, user_id int
	var registered_at time.Time

	rows, err := pool.Query(context.Background(),
		"SELECT ep.id, events.id, events.name, users.id, concat(users.name,' ',users.surname), ep.participants_status, ep.registered_at"+
			" FROM event_participants ep"+
			" JOIN events ON ep.event_id = events.id"+
			" JOIN users ON ep.user_id = users.id")

	var participants []model.EventParticipant
	for rows.Next() {
		if err := rows.Scan(&id, &event_id, &event_name, &user_id, &user_name, &participants_status, &registered_at); err != nil {
			return nil, err
		}
		participant := model.NewEventParticipant(uint(id), event_id, event_name, user_id, user_name, participants_status, registered_at)
		participants = append(participants, *participant)
	}
	rows.Close()
	fmt.Println(err)
	return participants, nil
}

func SelectOneEventParticipant(pool *pgxpool.Pool, ids int) (*model.EventParticipant, error) {

	var event_name, user_name string
	var participants_status model.ParticipantStatus
	var id, event_id, user_id int
	var registered_at time.Time

	row := pool.QueryRow(context.Background(),
		"SELECT ep.id, events.id, events.name, users.id, concat(users.name,' ',users.surname), ep.participants_status, ep.registered_at"+
			" FROM event_participants ep"+
			" JOIN events ON ep.event_id = events.id"+
			" JOIN users ON ep.user_id = users.id"+
			" WHERE ep.id = $1", ids)

	if err := row.Scan(&id, &event_id, &event_name, &user_id, &user_name, &participants_status, &registered_at); err != nil {
		return nil, err
	}
	participant := model.NewEventParticipant(uint(id), event_id, event_name, user_id, user_name, participants_status, registered_at)
	return participant, nil
}

// Вывод всех участников конкретного мероприятия
func SelectParticipantsByEventId(pool *pgxpool.Pool, event_id int) ([]model.EventParticipant, error) {

	var event_name, user_name string
	var participants_status model.ParticipantStatus
	var id, user_id int
	var registered_at time.Time

	rows, err := pool.Query(context.Background(),
		"SELECT ep.id, events.id, events.name, users.id, concat(users.name,' ',users.surname), ep.participants_status, ep.registered_at"+
			" FROM event_participants ep"+
			" JOIN events ON ep.event_id = events.id"+
			" JOIN users ON ep.user_id = users.id"+
			" WHERE ep.event_id = $1", event_id)

	var participants []model.EventParticipant
	for rows.Next() {
		if err := rows.Scan(&id, &event_id, &event_name, &user_id, &user_name, &participants_status, &registered_at); err != nil {
			return nil, err
		}
		participant := model.NewEventParticipant(uint(id), event_id, event_name, user_id, user_name, participants_status, registered_at)
		participants = append(participants, *participant)
	}
	rows.Close()
	fmt.Println(err)
	return participants, nil
}

// Изменение участника
func UpdateOneEventParticipant(pool *pgxpool.Pool, ids int, participant *model.EventParticipant) int64 {

	result, err := pool.Exec(context.Background(),
		"UPDATE event_participants SET event_id=$1, user_id=$2, participants_status=$3 WHERE id = $4",
		participant.GetEvent_id(),
		participant.GetUser_id(),
		participant.GetParticipants_status(),
		ids)
	if err != nil {
		return 0
	}
	return result.RowsAffected()
}

// Изменение участника
func UpdateOneEventStatus_OKe(pool *pgxpool.Pool, ids int) int64 {

	// Получаем текущий статус
	var current_status model.ParticipantStatus
	row := pool.QueryRow(context.Background(),
		"SELECT participants_status FROM event_participants WHERE id = $1", ids)

	if err := row.Scan(&current_status); err != nil {
		return 0
	}

	// Переключаем статус
	var new_status model.ParticipantStatus
	if current_status == model.StatusCame {
		new_status = model.StatusAbsent
	} else {
		new_status = model.StatusCame
	}

	// Обновляем
	result, err := pool.Exec(context.Background(),
		"UPDATE event_participants SET participants_status=$1 WHERE id = $2",
		new_status,
		ids)
	if err != nil {
		return 0
	}
	return result.RowsAffected()
}

func DeleteOneEventParticipant(pool *pgxpool.Pool, ids int) (int64, error) {

	result, err := pool.Exec(context.Background(), "DELETE FROM event_participants WHERE id = $1", ids)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), err
}
