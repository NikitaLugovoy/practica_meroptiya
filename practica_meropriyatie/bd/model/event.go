package model

import "time"

type CategoryEvents string
type EventStatus string

const (
	CategoryConference   CategoryEvents = "КОНФЕРЕНЦИЯ"
	CategorySubbotnik    CategoryEvents = "СУББОТНИК"
	CategoryOlympiad     CategoryEvents = "ОЛИМПИАДА"
	CategoryProfessional CategoryEvents = "ПРОФЕССИОНАЛЬНАЯ"
	CategoryTraining     CategoryEvents = "ТРЕНИНГ"
	CategorySport        CategoryEvents = "СПОРТИВНАЯ"
)

const (
	StatusPlanned   EventStatus = "ЗАПЛАНИРОВАНО"
	StatusCompleted EventStatus = "ПРОВЕДЕНО"
)

type Event struct {
	Id               uint           `json:"id"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	Date_time        time.Time      `json:"date_time"`
	Location         string         `json:"location"`
	Category         CategoryEvents `json:"category_events"`
	Status           EventStatus    `json:"status"`
	Organizer_id     int            `json:"organizer_id"`
	Organizer_name   string         `json:"organizer_name"`
	Responsible_id   int            `json:"responsible_id"`
	Responsible_name string         `json:"responsible_name"`
	CreatedAt        time.Time      `json:"created_at"`
}

func (e *Event) SetId(id uint)                       { e.Id = id }
func (e *Event) SetName(name string)                 { e.Name = name }
func (e *Event) SetDescription(description string)   { e.Description = description }
func (e *Event) SetDate_time(date_time time.Time)    { e.Date_time = date_time }
func (e *Event) SetLocation(location string)         { e.Location = location }
func (e *Event) SetCategory(category CategoryEvents) { e.Category = category }
func (e *Event) SetStatus(status EventStatus)        { e.Status = status }
func (e *Event) SetOrganizer_id(id int)              { e.Organizer_id = id }
func (e *Event) SetOrganizer_name(name string)       { e.Organizer_name = name }
func (e *Event) SetResponsible_id(id int)            { e.Responsible_id = id }
func (e *Event) SetResponsible_name(name string)     { e.Responsible_name = name }
func (e *Event) SetCreatedAt(createdAt time.Time)    { e.CreatedAt = createdAt }

func (e *Event) GetId() uint                 { return e.Id }
func (e *Event) GetName() string             { return e.Name }
func (e *Event) GetDescription() string      { return e.Description }
func (e *Event) GetDate_time() time.Time     { return e.Date_time }
func (e *Event) GetLocation() string         { return e.Location }
func (e *Event) GetCategory() CategoryEvents { return e.Category }
func (e *Event) GetStatus() EventStatus      { return e.Status }
func (e *Event) GetOrganizer_id() int        { return e.Organizer_id }
func (e *Event) GetOrganizer_name() string   { return e.Organizer_name }
func (e *Event) GetResponsible_id() int      { return e.Responsible_id }
func (e *Event) GetResponsible_name() string { return e.Responsible_name }
func (e *Event) GetCreatedAt() time.Time     { return e.CreatedAt }

func NewEvent(id uint, name string, description string, date_time time.Time, location string,
	category CategoryEvents, status EventStatus,
	organizer_id int, organizer_name string,
	responsible_id int, responsible_name string,
	createdAt time.Time) *Event {
	return &Event{
		Id:               id,
		Name:             name,
		Description:      description,
		Date_time:        date_time,
		Location:         location,
		Category:         category,
		Status:           status,
		Organizer_id:     organizer_id,
		Organizer_name:   organizer_name,
		Responsible_id:   responsible_id,
		Responsible_name: responsible_name,
		CreatedAt:        createdAt,
	}
}
