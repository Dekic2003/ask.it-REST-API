package questions

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"main/db"
	"main/utils"
	"net/http"
)



func GetByAuthorId(w http.ResponseWriter, r *http.Request) {

	params:=mux.Vars(r)
	id:=params["id"]

	var questions []Question
	results, err := db.Connection.Query("SELECT * FROM questions WHERE author_id=? ",id)
	if err!=nil {
		panic(err)
	}
	for results.Next(){
		var question Question
		err=results.Scan(&question.Id,&question.AuthorId,&question.Question,&question.CreatedAt,&question.UpdatedAt)
		if err!=nil{
			panic(err)
		}
		questions=append(questions,question)

	}
	utils.WriteSuccess(w,questions,true)


}
