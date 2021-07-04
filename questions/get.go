package questions

import (
	_ "github.com/go-sql-driver/mysql"
	"main/db"
	"main/utils"
	"net/http"
)



func Get(w http.ResponseWriter, r *http.Request) {

	var questions []Question
	results, err := db.Connection.Query("SELECT questions.id, users.id as author_id, users.email, questions.question, (SELECT  COUNT(*) FROM question_reactions WHERE question_id=questions.id AND rating=true) as likes, (SELECT  COUNT(*) FROM question_reactions WHERE question_id=questions.id AND rating=false) as dislikes, questions.created_at, questions.updated_at FROM questions INNER JOIN users on questions.author_id = users.id")
	if err!=nil {
		utils.WriteError(w,"Unable to fetch questions",err,http.StatusInternalServerError)
		return	}
	for results.Next(){
		var question Question
		err=results.Scan(&question.Id,&question.AuthorId,&question.Author,&question.Question,&question.Likes,&question.Dislikes,&question.CreatedAt,&question.UpdatedAt)
		if err!=nil{
			utils.WriteError(w,"Unable to scan results",err,http.StatusInternalServerError)
			return
		}
		questions=append(questions,question)

	}
	utils.WriteSuccess(w,questions,true)


}

