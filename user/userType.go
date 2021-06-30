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
type ChangePasswordUser struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	NewPassword string `json:"newPassword"`
}
type ChangeUserInfo struct {
	Id int `json:"id"`
	Name string`json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
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
type Notification struct {
	Id int `json:"id"`
	QuestionId int `json:"question_id"`
	QuestionAuthorId int `json:"question_author_id"`
	AnswerAuthorId int `json:"answer_author_id"`
	AuthorAnswerName string `json:"author_answer_name"`
	IsRead bool `json:"isRead"`
}
type NotificationUpdate struct {
	Id int `json:"id"`
}
type returnUserById struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
}