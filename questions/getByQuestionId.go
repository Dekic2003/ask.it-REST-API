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
	err := db.Connection.QueryRow("SELECT * FROM question WHERE id=?",id).Scan(&question.Id,&question.AuthorId,&question.Question,&question.Likes,&question.Dislikes,&question.CreatedAt,&question.UpdatedAt)
	if err!=nil {
		utils.WriteError(w,"Unable to fetch question",err,http.StatusInternalServerError)
	}

	utils.WriteSuccess(w,question)

}
