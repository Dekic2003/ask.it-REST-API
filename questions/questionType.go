package questions

import "time"

type EditedQuestion struct {
	Id int `json:"id"`
	Question string `json:"question"`
}

type Question struct {
	Id string `json:"id"`
	AuthorId string `json:"author_id"`
	Question string `json:"question"`
	Likes int `json:"likes"`
	Dislikes int `json:"dislikes"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
type NewQuestion struct {
	AuthorId int `json:"author_id"`
	Question string `json:"question"`
}
type DeleteQuestion struct {
	Id int `json:"id"`
}