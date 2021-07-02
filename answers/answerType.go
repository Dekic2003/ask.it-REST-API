package answers

import "time"

type Answer struct {
	Id int `json:"id"`
	QuestionID int `json:"question_id"`
	AuthorId int `json:"author_id"`
	Author string `json:"author"`
	QuestionAuthorId int `json:"question_author_id"`
	Answer string `json:"answer"`
	Likes int `json:"likes"`
	Dislikes int `json:"dislikes"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
type NewAnswer struct {
	QuestionID int `json:"question_id"`
	QuestionAuthorId int `json:"question_author_id"`
	AuthorId int `json:"author_id"`
	Answer string `json:"answer"`

}
type EditedAnswer struct {
	Id int `json:"id"`
	Answer string `json:"answer"`
}
type DeleteAnswer struct {
	Id int `json:"id"`
}
type AnswerReaction struct {
	AuthorId int `json:"author_id"`
	AnswerId int `json:"answer_id"`
	Reaction bool `json:"reaction"`
}