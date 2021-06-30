package questions

import (
	_ "github.com/go-sql-driver/mysql"
	"main/db"
	"main/utils"
	"net/http"
)



func GetHotQuestions(w http.ResponseWriter, r *http.Request) {

	var questions []Question
	results, err := db.Connection.Query("SELECT question.id, users.id as author_id, users.email, question.question, (SELECT  COUNT(*) FROM question_reactions WHERE question_id=question.id AND rating=true) as likes, (SELECT  COUNT(*) FROM question_reactions WHERE question_id=question.id AND rating=false) as dislikes, question.created_at, question.updated_at FROM question INNER JOIN users on question.author_id = users.id ORDER BY likes DESC ")
	if err!=nil {
	utils.WriteError(w,"Unable to get hot questions",err,http.StatusInternalServerError)
	}
	for results.Next(){
		var question Question
		err=results.Scan(&question.Id,&question.AuthorId,&question.Author,&question.Question,&question.Likes,&question.Dislikes,&question.CreatedAt,&question.UpdatedAt)
		if err!=nil{
			utils.WriteError(w,"Unable to scan results",err,http.StatusInternalServerError)
		}
		questions=append(questions,question)

	}
	utils.WriteSuccess(w,questions,true)


}

