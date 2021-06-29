package questions

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"main/db"
	"net/http"
)



func Get(w http.ResponseWriter, r *http.Request) {

	var questions []Question
	results, err := db.Connection.Query("SELECT question.id, users.id as author_id, users.email, question.question, question.likes, question.dislikes, question.created_at, question.updated_at FROM question INNER JOIN users on question.author_id = users.id")
	if err!=nil {
		panic(err)
	}
	for results.Next(){
		var question Question
		err=results.Scan(&question.Id,&question.AuthorId,&question.Author,&question.Question,&question.Likes,&question.Dislikes,&question.CreatedAt,&question.UpdatedAt)
		if err!=nil{
			panic(err)
		}
		questions=append(questions,question)

	}
	res,err := json.Marshal(questions)
	w.Header().Set("Content-Type","application/json")
	w.Write(res)


}

