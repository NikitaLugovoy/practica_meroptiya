package model

import (
	"time"
)

type User struct {
	Id           uint      `json:"id"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Login        string    `json:"login"`
	Password     string    `json:"password"`
	Phone_number string    `json:"phone_number"`
	Email        string    `json:"email"`
	Url_avatar   string    `json:"url_avatar"`
	User_role    string    `json:"user_role"`
	CreatedAt    time.Time `json:"created_at"`
}

func (user *User) SetId(id uint) {
	user.Id = id
}

func (user *User) SetName(name string) {
	user.Name = name
}

func (user *User) SetSurname(surname string) {
	user.Surname = surname
}

func (user *User) SetLogin(login string) {
	user.Login = login
}

func (user *User) SetPassword(password string) {
	user.Password = password
}

func (user *User) SetPhoneNumber(phone_number string) {
	user.Phone_number = phone_number
}

func (user *User) SetEmail(email string) {
	user.Email = email
}

func (user *User) SetUrl_avatar(url_avatar string) {
	user.Url_avatar = url_avatar
}

func (user *User) SetUser_role(user_role string) {
	user.User_role = user_role
}

func (user *User) SetCreatedAt(created_at time.Time) {
	user.CreatedAt = created_at
}

func (user *User) GetId() uint {
	return user.Id
}
func (user *User) GetName() string {
	return user.Name
}
func (user *User) GetSurname() string {
	return user.Surname
}
func (user *User) GetLogin() string {
	return user.Login
}
func (user *User) GetPassword() string {
	return user.Password
}
func (user *User) GetPhoneNumber() string {
	return user.Phone_number
}
func (user *User) GetEmail() string {
	return user.Email
}
func (user *User) GetUrl_avatar() string {
	return user.Url_avatar
}
func (user *User) GetUser_role() string {
	return user.User_role
}
func (user *User) GetCreatedAt() time.Time {
	return user.CreatedAt
}

func NewUser(id uint, name string, surname string, login string, password string, phone_number string, email string, url_avatar string, user_role string, created_at time.Time) *User {
	return &User{
		Id:           id,
		Name:         name,
		Surname:      surname,
		Login:        login,
		Password:     password,
		Phone_number: phone_number,
		Email:        email,
		Url_avatar:   url_avatar,
		User_role:    user_role,
		CreatedAt:    created_at,
	}
}
