package questions

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"main/db"
	"main/utils"
	"net/http"
)



func GetByQuestionId(w http.ResponseWriter, r *http.Request) {

	params:=mux.Vars(r)
	id:=params["id"]
	var question Question
	err := db.Connection.QueryRow("SELECT id, author_id,(SELECT email FROM users WHERE users.id=author_id) as author, question, (SELECT  COUNT(*) FROM question_reactions WHERE question_id=question.id AND rating=true) as likes, (SELECT  COUNT(*) FROM question_reactions WHERE question_id=question.id AND rating=false) as dislikes, created_at, updated_at FROM question WHERE id=?",id).Scan(&question.Id,&question.AuthorId,&question.Author,&question.Question,&question.Likes,&question.Dislikes,&question.CreatedAt,&question.UpdatedAt)
	if err!=nil {
		utils.WriteError(w,"Unable to fetch question",err,http.StatusInternalServerError)
		return
	}

	utils.WriteSuccess(w,question,true)

}
