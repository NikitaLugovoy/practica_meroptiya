package model

import "time"

type EventImage struct {
	Id        uint      `json:"id"`
	Event_id  int       `json:"event_id"`
	Url_image string    `json:"url_image"`
	CreatedAt time.Time `json:"created_at"`
}

func (ei *EventImage) SetId(id uint)            { ei.Id = id }
func (ei *EventImage) SetEvent_id(event_id int) { ei.Event_id = event_id }
func (ei *EventImage) SetUrl_image(url string)  { ei.Url_image = url }
func (ei *EventImage) SetCreatedAt(t time.Time) { ei.CreatedAt = t }

func (ei *EventImage) GetId() uint             { return ei.Id }
func (ei *EventImage) GetEvent_id() int        { return ei.Event_id }
func (ei *EventImage) GetUrl_image() string    { return ei.Url_image }
func (ei *EventImage) GetCreatedAt() time.Time { return ei.CreatedAt }

func NewEventImage(id uint, event_id int, url_image string, createdAt time.Time) *EventImage {
	return &EventImage{
		Id:        id,
		Event_id:  event_id,
		Url_image: url_image,
		CreatedAt: createdAt,
	}
}
