// model/event_participant.go
package model

import "time"

type ParticipantStatus string

const (
	StatusCame   ParticipantStatus = "ПРИШЁЛ"
	StatusAbsent ParticipantStatus = "ОТСУТСТВОВАЛ"
)

type EventParticipant struct {
	Id                  uint              `json:"id"`
	Event_id            int               `json:"event_id"`
	Event_name          string            `json:"event_name"`
	User_id             int               `json:"user_id"`
	User_name           string            `json:"user_name"`
	Participants_status ParticipantStatus `json:"participants_status"`
	RegisteredAt        time.Time         `json:"registered_at"`
}

func (ep *EventParticipant) SetId(id uint)                              { ep.Id = id }
func (ep *EventParticipant) SetEvent_id(event_id int)                   { ep.Event_id = event_id }
func (ep *EventParticipant) SetEvent_name(name string)                  { ep.Event_name = name }
func (ep *EventParticipant) SetUser_id(user_id int)                     { ep.User_id = user_id }
func (ep *EventParticipant) SetUser_name(name string)                   { ep.User_name = name }
func (ep *EventParticipant) SetParticipants_status(s ParticipantStatus) { ep.Participants_status = s }
func (ep *EventParticipant) SetRegisteredAt(t time.Time)                { ep.RegisteredAt = t }

func (ep *EventParticipant) GetId() uint                               { return ep.Id }
func (ep *EventParticipant) GetEvent_id() int                          { return ep.Event_id }
func (ep *EventParticipant) GetEvent_name() string                     { return ep.Event_name }
func (ep *EventParticipant) GetUser_id() int                           { return ep.User_id }
func (ep *EventParticipant) GetUser_name() string                      { return ep.User_name }
func (ep *EventParticipant) GetParticipants_status() ParticipantStatus { return ep.Participants_status }
func (ep *EventParticipant) GetRegisteredAt() time.Time                { return ep.RegisteredAt }

func NewEventParticipant(id uint, event_id int, event_name string,
	user_id int, user_name string, participants_status ParticipantStatus,
	registeredAt time.Time) *EventParticipant {
	return &EventParticipant{
		Id:                  id,
		Event_id:            event_id,
		Event_name:          event_name,
		User_id:             user_id,
		User_name:           user_name,
		Participants_status: participants_status,
		RegisteredAt:        registeredAt,
	}
}
