package model

import "time"

type EventResult struct {
	Id        uint      `json:"id"`
	Event_id  int       `json:"event_id"`
	Result    string    `json:"result"`
	CreatedAt time.Time `json:"created_at"`
}

func (er *EventResult) SetId(id uint)            { er.Id = id }
func (er *EventResult) SetEvent_id(event_id int) { er.Event_id = event_id }
func (er *EventResult) SetResult(result string)  { er.Result = result }
func (er *EventResult) SetCreatedAt(t time.Time) { er.CreatedAt = t }

func (er *EventResult) GetId() uint             { return er.Id }
func (er *EventResult) GetEvent_id() int        { return er.Event_id }
func (er *EventResult) GetResult() string       { return er.Result }
func (er *EventResult) GetCreatedAt() time.Time { return er.CreatedAt }

func NewEventResult(id uint, event_id int, result string, createdAt time.Time) *EventResult {
	return &EventResult{
		Id:        id,
		Event_id:  event_id,
		Result:    result,
		CreatedAt: createdAt,
	}
}
