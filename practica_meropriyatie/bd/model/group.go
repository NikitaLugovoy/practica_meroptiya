package model

import "time"

type Group struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (g *Group) SetId(id uint)            { g.Id = id }
func (g *Group) SetName(name string)      { g.Name = name }
func (g *Group) SetCreatedAt(t time.Time) { g.CreatedAt = t }

func (g *Group) GetId() uint             { return g.Id }
func (g *Group) GetName() string         { return g.Name }
func (g *Group) GetCreatedAt() time.Time { return g.CreatedAt }

func NewGroup(id uint, name string, createdAt time.Time) *Group {
	return &Group{
		Id:        id,
		Name:      name,
		CreatedAt: createdAt,
	}
}
