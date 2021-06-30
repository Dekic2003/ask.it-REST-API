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
	results, err := db.Connection.Query("SELECT * FROM answer WHERE question_id=?",id)
	if err!=nil {
		utils.WriteError(w,"Unable to fetch question",err,http.StatusInternalServerError)
		return
	}
	for results.Next(){
		var answer Answer
		err=results.Scan(&answer.Id,&answer.QuestionID,&answer.AuthorId,&answer.Answer,&answer.Likes,&answer.Dislikes,&answer.CreatedAt,&answer.UpdatedAt)
		if err!=nil{
			utils.WriteError(w,"Unable to scan results",err,http.StatusInternalServerError)
			return
		}
		answers=append(answers,answer)

	}
	utils.WriteSuccess(w,"",true)



}
