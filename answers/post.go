package answers

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"main/db"
	"main/utils"
	"net/http"
)


func Post(w http.ResponseWriter, r *http.Request) {
	//userId:=r.Header.Get("user-id")
	req,err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w,"Unable to update",err,http.StatusInternalServerError)
		return
	}
	var answer NewAnswer
	json.Unmarshal(req,&answer)
	_,err=db.Connection.Exec("INSERT INTO answer(question_id, author_id, answer) VALUES (?,?,?)",answer.QuestionID,answer.AuthorId,answer.Answer)
	if err != nil {
		utils.WriteError(w,"Unable to post answer",err,http.StatusInternalServerError)
		return
	}
	_,err=db.Connection.Exec("INSERT INTO notifications (question_id, question_author_id, answer_author_id) VALUES (?,?,?);",answer.QuestionID,answer.QuestionAuthorId,answer.AuthorId)
	if err != nil {
		utils.WriteError(w,"Unable to send notification",err,http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w,"",true)
}


