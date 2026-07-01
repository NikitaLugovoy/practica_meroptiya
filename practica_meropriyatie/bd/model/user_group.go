package model

import "time"

type UserGroup struct {
	Id         uint      `json:"id"`
	User_id    int       `json:"user_id"`
	User_name  string    `json:"user_name"`
	Group_id   int       `json:"group_id"`
	Group_name string    `json:"group_name"`
	CreatedAt  time.Time `json:"created_at"`
}

func (ug *UserGroup) SetId(id uint)             { ug.Id = id }
func (ug *UserGroup) SetUser_id(user_id int)    { ug.User_id = user_id }
func (ug *UserGroup) SetUser_name(name string)  { ug.User_name = name }
func (ug *UserGroup) SetGroup_id(group_id int)  { ug.Group_id = group_id }
func (ug *UserGroup) SetGroup_name(name string) { ug.Group_name = name }
func (ug *UserGroup) SetCreatedAt(t time.Time)  { ug.CreatedAt = t }

func (ug *UserGroup) GetId() uint             { return ug.Id }
func (ug *UserGroup) GetUser_id() int         { return ug.User_id }
func (ug *UserGroup) GetUser_name() string    { return ug.User_name }
func (ug *UserGroup) GetGroup_id() int        { return ug.Group_id }
func (ug *UserGroup) GetGroup_name() string   { return ug.Group_name }
func (ug *UserGroup) GetCreatedAt() time.Time { return ug.CreatedAt }

func NewUserGroup(id uint, user_id int, user_name string,
	group_id int, group_name string, createdAt time.Time) *UserGroup {
	return &UserGroup{
		Id:         id,
		User_id:    user_id,
		User_name:  user_name,
		Group_id:   group_id,
		Group_name: group_name,
		CreatedAt:  createdAt,
	}
}
