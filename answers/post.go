package answers

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"main/db"
	"net/http"
)

type DefaultResp struct {
	Success bool `json:"success"`
}

func Post(w http.ResponseWriter, r *http.Request) {

	req,err := ioutil.ReadAll(r.Body)
	if err != nil{
		panic(err)
	}
	resp:=DefaultResp{Success: true}
	var answer NewAnswer
	json.Unmarshal(req,&answer)
	db.Connection.Query("INSERT INTO answer(question_id, author_id, answer) VALUES (?,?,?)",answer.QuestionID,answer.AuthorId,answer.Answer)
	res,err := json.Marshal(resp)
	w.Header().Set("Content-Type","application/json")
	w.Write(res)
}


