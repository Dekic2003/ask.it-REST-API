package answers

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"main/db"
	"net/http"
)



func Get(w http.ResponseWriter, r *http.Request) {

	params:=mux.Vars(r)
	id:=params["id"]
 	var answers []Answer
	results, err := db.Connection.Query("SELECT * FROM answer WHERE question_id=?",id)
	if err!=nil {
		panic(err)
	}
	for results.Next(){
		var answer Answer
		err=results.Scan(&answer.Id,&answer.QuestionID,&answer.AuthorId,&answer.Answer,&answer.Likes,&answer.Dislikes,&answer.CreatedAt,&answer.UpdatedAt)
		if err!=nil{
			panic(err)
		}
		answers=append(answers,answer)

	}
	res,err := json.Marshal(answers)
	w.Header().Set("Content-Type","application/json")
	w.Write(res)


}
