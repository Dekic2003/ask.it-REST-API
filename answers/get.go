package answers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"main/db"
	"main/utils"
	"net/http"
)



func Get(w http.ResponseWriter, r *http.Request) {

	params:=mux.Vars(r)
	id:=params["id"]
 	var answers []Answer
	results, err := db.Connection.Query("SELECT id, question_id, author_id, (SELECT email FROM users WHERE users.id=author_id) as author,answer, (SELECT  COUNT(*) FROM answer_reactions WHERE answer_id=answers.id AND rating=true) as likes, (SELECT  COUNT(*) FROM answer_reactions WHERE answer_id=answers.id AND rating=false) as dislikes, created_at, updated_at FROM answers WHERE question_id=?",id)
	if err!=nil {
		utils.WriteError(w,"Unable to fetch question",err,http.StatusInternalServerError)
		return
	}
	for results.Next(){
		var answer Answer
		err=results.Scan(&answer.Id,&answer.QuestionID,&answer.AuthorId,&answer.Author,&answer.Answer,&answer.Likes,&answer.Dislikes,&answer.CreatedAt,&answer.UpdatedAt)
		if err!=nil{
			utils.WriteError(w,"Unable to scan results",err,http.StatusInternalServerError)
			return
		}
		answers=append(answers,answer)

	}
	utils.WriteSuccess(w,answers,true)



}
