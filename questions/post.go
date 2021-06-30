package questions

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"main/db"
	"main/utils"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {

	req,err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, "Unable to read body", err, http.StatusInternalServerError)
		return
	}
	var question NewQuestion
	json.Unmarshal(req,&question)
	_,err =db.Connection.Exec("INSERT INTO question(author_id,question) VALUES (?,?) ",question.AuthorId,question.Question)
	if err!=nil{
		utils.WriteError(w,"Unable to post question",err,http.StatusInternalServerError)
		return
	}
	utils.WriteSuccess(w,"",true)

}

