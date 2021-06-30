package questions

import "time"

type EditedQuestion struct {
	Id int `json:"id"`
	Question string `json:"question"`
}

type Question struct {
	Id int `json:"id"`
	AuthorId int `json:"author_id"`
	Author string `json:"author"`
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
type QuestionReaction struct {
	AuthorId int `json:"author_id"`
	QuestionId int `json:"question_id"`
	Reaction bool `json:"reaction"`
}