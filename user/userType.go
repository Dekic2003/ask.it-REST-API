package user

import "time"

type NewUser struct {
	Name string`json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Password string `json:"password"`
}
type LoginUser struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
type DbUser struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
type returnUser struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	AccessToken string `json:"access_token"`
}
type LeaderboardUser struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Answers int `json:"answers"`
	Questions int `json:"questions"`
}