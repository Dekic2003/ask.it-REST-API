package questions

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"main/db"
	"net/http"
)



func Get(w http.ResponseWriter, r *http.Request) {

	var questions []Question
	results, err := db.Connection.Query("SELECT * FROM question")
	if err!=nil {
		panic(err)
	}
	for results.Next(){
		var question Question
		err=results.Scan(&question.Id,&question.AuthorId,&question.Question,&question.Likes,&question.Dislikes,&question.CreatedAt,&question.UpdatedAt)
		if err!=nil{
			panic(err)
		}
		questions=append(questions,question)

	}
	res,err := json.Marshal(questions)
	w.Header().Set("Content-Type","application/json")
	w.Write(res)


}

